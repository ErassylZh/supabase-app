package worker

import (
	"context"
	"log"
	"strings"
	"time"
	"work-project/internal/airtable"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type AirTableSync struct {
	airTable       repository.AirTable
	product        repository.Product
	post           repository.Post
	storage        repository.StorageClient
	image          repository.Image
	hashtag        repository.Hashtag
	postHashtag    repository.PostHashtag
	collection     repository.Collection
	postCollection repository.PostCollection
	stories        repository.Stories
	storyPage      repository.StoryPage
}

func NewAirTableSync(
	airTable repository.AirTable,
	product repository.Product,
	post repository.Post,
	storage repository.StorageClient,
	image repository.Image,
	hashtag repository.Hashtag,
	postHashtag repository.PostHashtag,
	collection repository.Collection,
	postCollection repository.PostCollection,
	stories repository.Stories,
	storyPage repository.StoryPage,
) *AirTableSync {
	return &AirTableSync{
		airTable:       airTable,
		product:        product,
		post:           post,
		storage:        storage,
		image:          image,
		hashtag:        hashtag,
		postHashtag:    postHashtag,
		collection:     collection,
		postCollection: postCollection,
		stories:        stories,
		storyPage:      storyPage,
	}
}

func (h *AirTableSync) Run() (err error) {
	ctx := context.Background()
	if err := h.syncHashtags(ctx); err != nil {
		log.Println("error while syncing hashtags:", err)
		return err
	}

	if err := h.syncCollections(ctx); err != nil {
		log.Println("error while syncing collection:", err)
		return err
	}

	if err := h.syncProducts(ctx); err != nil {
		log.Println("error while syncing products:", err)
		return err
	}

	if err := h.syncPosts(ctx); err != nil {
		log.Println("error while syncing posts:", err)
		return err
	}

	if err := h.syncStories(ctx); err != nil {
		log.Println("error while syncing stories:", err)
		return err
	}

	log.Println("airtable sync completed successfully")
	return nil
}

func (h *AirTableSync) syncProducts(ctx context.Context) error {
	products, err := h.airTable.GetProducts(ctx)
	if err != nil {
		return err
	}
	productsAirtableBySku := make(map[string]airtable.BaseObject[airtable.ProductListResponse])
	for _, product := range products {
		productsAirtableBySku[product.Fields.Sku] = product
	}

	productsDb, err := h.product.GetAll(ctx)
	if err != nil {
		return err
	}
	productsDbBySku := make(map[string]model.Product)
	for _, product := range productsDb {
		productsDbBySku[product.Sku] = product
	}

	newProducts := make([]model.Product, 0)
	updateProducts := make([]model.Product, 0)
	for sku := range productsAirtableBySku {
		if product, exists := productsDbBySku[sku]; exists {
			if product.Point != productsAirtableBySku[sku].Fields.Point ||
				product.Count != productsAirtableBySku[sku].Fields.Count ||
				product.Description != productsAirtableBySku[sku].Fields.Description ||
				product.Title != productsAirtableBySku[sku].Fields.Title ||
				product.Sapphire != productsAirtableBySku[sku].Fields.Sapphire ||
				product.SellType != productsAirtableBySku[sku].Fields.SellType ||
				product.ProductType != productsAirtableBySku[sku].Fields.ProductType ||
				product.Status != productsAirtableBySku[sku].Fields.Status {

				product.Point = productsAirtableBySku[sku].Fields.Point
				product.Sapphire = productsAirtableBySku[sku].Fields.Sapphire
				product.Count = productsAirtableBySku[sku].Fields.Count
				product.Description = productsAirtableBySku[sku].Fields.Description
				product.Title = productsAirtableBySku[sku].Fields.Title
				product.SellType = productsAirtableBySku[sku].Fields.SellType
				product.ProductType = productsAirtableBySku[sku].Fields.ProductType
				product.Status = productsAirtableBySku[sku].Fields.Status
				updateProducts = append(updateProducts, product)
			}
			continue
		}
		newProducts = append(newProducts, model.Product{
			Title:             productsAirtableBySku[sku].Fields.Title,
			Sku:               productsAirtableBySku[sku].Fields.Sku,
			Description:       productsAirtableBySku[sku].Fields.Description,
			Point:             productsAirtableBySku[sku].Fields.Point,
			Sapphire:          productsAirtableBySku[sku].Fields.Sapphire,
			Status:            productsAirtableBySku[sku].Fields.Status,
			Count:             productsAirtableBySku[sku].Fields.Count,
			AirtableProductId: productsAirtableBySku[sku].Id,
			SellType:          productsAirtableBySku[sku].Fields.SellType,
			ProductType:       productsAirtableBySku[sku].Fields.ProductType,
		})
	}
	if len(newProducts) > 0 {
		newProducts, err = h.product.CreateMany(ctx, newProducts)
		if err != nil {
			log.Println(ctx, "error while create new products from airtable ", "err", err)
			return err
		}

		imagesProduct := make([]model.Image, 0)
		for _, np := range newProducts {
			productId := np.ProductID
			for _, img := range productsAirtableBySku[np.Sku].Fields.Image {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_PRODUCT), img.FileName, img.Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "pr name", np.Title)
					return err
				}
				log.Println(ctx, "file for "+np.Title+" saved")
				imagesProduct = append(imagesProduct, model.Image{
					ProductID: &productId,
					ImageUrl:  file,
					FileName:  img.FileName,
				})
			}
		}
		_, err = h.image.CreateMany(ctx, imagesProduct)
		if err != nil {
			log.Println(ctx, "error while create images from airtable ", "err", err)
			return err
		}
	}

	if len(updateProducts) > 0 {
		_, err = h.product.UpdateMany(ctx, updateProducts)
		if err != nil {
			log.Println(ctx, "error while update exist products from airtable ", "err", err)
			return err
		}
	}
	return nil
}

func (h *AirTableSync) syncPosts(ctx context.Context) error {
	posts, err := h.airTable.GetPosts(ctx)
	if err != nil {
		return err
	}

	postsAirtableByUuid := make(map[string]airtable.BaseObject[airtable.Post])
	for _, post := range posts {
		if post.Fields.Company == nil {
			continue
		}
		postsAirtableByUuid[post.Fields.Uuid] = post
	}

	postsDb, err := h.post.GetAll(ctx)
	if err != nil {
		return err
	}

	postsDbByUuid := make(map[string]model.Post)
	for _, post := range postsDb {
		postsDbByUuid[post.Uuid] = post
	}

	newPosts := make([]model.Post, 0)
	updatePosts := make([]model.Post, 0)
	for uuid := range postsAirtableByUuid {
		if post, exists := postsDbByUuid[uuid]; exists {
			var existsHashtags []string
			for _, ht := range post.Hashtags {
				existsHashtags = append(existsHashtags, ht.Name)
			}
			var existsCollections []string
			for _, ht := range post.Collections {
				existsCollections = append(existsCollections, ht.Name)
			}
			if post.Company != postsAirtableByUuid[uuid].Fields.Company ||
				post.Language != postsAirtableByUuid[uuid].Fields.Language ||
				post.Title != postsAirtableByUuid[uuid].Fields.Title ||
				post.Description != postsAirtableByUuid[uuid].Fields.Description ||
				post.Status != postsAirtableByUuid[uuid].Fields.Status ||
				post.Body != postsAirtableByUuid[uuid].Fields.Body ||
				post.ReadTime != postsAirtableByUuid[uuid].Fields.ReadTime ||
				post.Point != postsAirtableByUuid[uuid].Fields.Point ||
				post.QuizTime != postsAirtableByUuid[uuid].Fields.QuizTime ||
				post.RatingStatus != postsAirtableByUuid[uuid].Fields.RatingStatus ||
				post.Sapphire != postsAirtableByUuid[uuid].Fields.Sapphire {

				post.Company = postsAirtableByUuid[uuid].Fields.Company
				post.Language = postsAirtableByUuid[uuid].Fields.Language
				post.Title = postsAirtableByUuid[uuid].Fields.Title
				post.Description = postsAirtableByUuid[uuid].Fields.Description
				post.Status = postsAirtableByUuid[uuid].Fields.Status
				post.Body = postsAirtableByUuid[uuid].Fields.Body
				post.ReadTime = postsAirtableByUuid[uuid].Fields.ReadTime
				post.Point = postsAirtableByUuid[uuid].Fields.Point
				post.QuizTime = postsAirtableByUuid[uuid].Fields.QuizTime
				post.RatingStatus = postsAirtableByUuid[uuid].Fields.RatingStatus
				post.Sapphire = postsAirtableByUuid[uuid].Fields.Sapphire
				updatePosts = append(updatePosts, post)
			}
			if !h.compareHashtags(existsHashtags, postsAirtableByUuid[uuid].Fields.HashtagName) {
				err = h.postHashtag.DeleteByPostId(ctx, post.PostID)
				if err != nil {
					return err
				}
				names := postsAirtableByUuid[uuid].Fields.HashtagName
				var postHashtags []model.PostHashtag
				for _, name := range names {
					ht, err := h.hashtag.GetByName(ctx, name)
					if err != nil {
						return err
					}
					postHashtags = append(postHashtags, model.PostHashtag{
						PostId:    post.PostID,
						HashtagId: ht.HashtagID,
					})
				}
				if len(postHashtags) > 0 {
					_, err = h.postHashtag.CreateMany(ctx, postHashtags)
					if err != nil {
						return err
					}
				}
			}
			if !h.compareHashtags(existsCollections, postsAirtableByUuid[uuid].Fields.CollectionName) {
				err = h.postCollection.DeleteByPostId(ctx, post.PostID)
				if err != nil {
					return err
				}
				names := postsAirtableByUuid[uuid].Fields.CollectionName
				var postCollections []model.PostCollection
				for _, name := range names {
					ht, err := h.collection.GetByName(ctx, name)
					if err != nil {
						return err
					}
					postCollections = append(postCollections, model.PostCollection{
						PostId:       post.PostID,
						CollectionId: ht.CollectionID,
					})
				}
				if len(postCollections) > 0 {
					_, err = h.postCollection.CreateMany(ctx, postCollections)
					if err != nil {
						return err
					}
				}
			}
			continue
		}

		newPosts = append(newPosts, model.Post{
			Company:      postsAirtableByUuid[uuid].Fields.Company,
			Language:     postsAirtableByUuid[uuid].Fields.Language,
			Title:        postsAirtableByUuid[uuid].Fields.Title,
			Uuid:         postsAirtableByUuid[uuid].Fields.Uuid,
			Description:  postsAirtableByUuid[uuid].Fields.Description,
			Status:       postsAirtableByUuid[uuid].Fields.Status,
			Body:         postsAirtableByUuid[uuid].Fields.Body,
			ReadTime:     postsAirtableByUuid[uuid].Fields.ReadTime,
			Point:        postsAirtableByUuid[uuid].Fields.Point,
			QuizTime:     postsAirtableByUuid[uuid].Fields.QuizTime,
			RatingStatus: postsAirtableByUuid[uuid].Fields.RatingStatus,
			Sapphire:     postsAirtableByUuid[uuid].Fields.Sapphire,
		})
	}

	if len(newPosts) > 0 {
		newPosts, err = h.post.CreateMany(ctx, newPosts)
		if err != nil {
			log.Println(ctx, "error while creating new posts from airtable:", err)
			return err
		}

		imagesProduct := make([]model.Image, 0)
		postHashtags := make([]model.PostHashtag, 0)
		for _, np := range newPosts {
			postId := np.PostID
			for _, hashtag := range postsAirtableByUuid[np.Uuid].Fields.HashtagName {
				hashtagObj, err := h.hashtag.GetByName(ctx, hashtag)
				if err != nil {
					return err
				}
				postHashtags = append(postHashtags, model.PostHashtag{
					PostId:    postId,
					HashtagId: hashtagObj.HashtagID,
				})
			}

			for _, img := range postsAirtableByUuid[np.Uuid].Fields.Image {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_POST), img.FileName, img.Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "pr name", np.Title)
				}
				log.Println(ctx, "image file for "+np.Title+" saved")
				imagesProduct = append(imagesProduct, model.Image{
					PostID:   &postId,
					ImageUrl: file,
					FileName: img.FileName,
					Type:     string(model.POST_IMAGE_TYPE_IMAGE),
				})
			}
			for _, img := range postsAirtableByUuid[np.Uuid].Fields.Logo {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_POST), img.FileName, img.Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "pr name", np.Title)
				}
				log.Println(ctx, "logo file for "+np.Title+" saved")
				imagesProduct = append(imagesProduct, model.Image{
					PostID:   &postId,
					ImageUrl: file,
					FileName: img.FileName,
					Type:     string(model.POST_IMAGE_TYPE_LOGO),
				})
			}
		}
		if len(postHashtags) > 0 {
			_, err = h.postHashtag.CreateMany(ctx, postHashtags)
			if err != nil {
				log.Println(ctx, "error while create post hashtags from airtable ", "err", err)
				return err
			}

		}
		_, err = h.image.CreateMany(ctx, imagesProduct)
		if err != nil {
			log.Println(ctx, "error while create images from airtable ", "err", err)
			return err
		}
	}

	if len(updatePosts) > 0 {
		_, err = h.post.UpdateMany(ctx, updatePosts)
		if err != nil {
			log.Println(ctx, "error while updating existing posts from airtable:", err)
			return err
		}
	}

	if len(postsAirtableByUuid) > 0 {
		uuids := make([]string, 0)
		for key := range postsAirtableByUuid {
			uuids = append(uuids, key)
		}
		err = h.post.DeleteAllNotInUuid(ctx, uuids)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *AirTableSync) syncStories(ctx context.Context) error {
	stories, err := h.airTable.GetStories(ctx)
	if err != nil {
		return err
	}

	imagesByUuidAirtable := make(map[string][]airtable.Image)
	storiesAirtableByTitle := make(map[string][]airtable.BaseObject[airtable.Stories])
	for _, story := range stories {
		if story.Fields.Status == string(model.POST_STATUS_DRAFT) {
			continue
		}
		storiesAirtableByTitle[story.Fields.Title] = []airtable.BaseObject[airtable.Stories]{}
		imagesByUuidAirtable[story.Fields.Uuid] = story.Fields.Image
	}
	for _, story := range stories {
		if story.Fields.Status == string(model.POST_STATUS_DRAFT) {
			continue
		}
		storiesAirtableByTitle[story.Fields.Title] = append(storiesAirtableByTitle[story.Fields.Title], story)
	}

	storiesDb, err := h.stories.GetAll(ctx)
	storiesDbByTitle := make(map[string]model.Stories)
	for _, story := range storiesDb {
		storiesDbByTitle[story.Title] = story
	}
	storyPagesByUuid := make(map[string]model.StoryPage)
	for _, story := range storiesDb {
		for _, sp := range story.StoryPages {
			storyPagesByUuid[sp.Uuid] = sp
		}
	}

	createStories := make([]model.Stories, 0)
	updateStories := make([]model.Stories, 0)
	updateStoryPages := make([]model.StoryPage, 0)
	for key := range storiesAirtableByTitle {
		if _, ok := storiesDbByTitle[key]; ok {
			for _, data := range storiesAirtableByTitle[key] {
				if data.Fields.Text != storyPagesByUuid[data.Fields.Uuid].Text ||
					data.Fields.Order != storyPagesByUuid[data.Fields.Uuid].PageOrder ||
					!strings.Contains(storyPagesByUuid[data.Fields.Uuid].ImagePath, data.Fields.Image[0].FileName) {
					updateStoryPages = append(updateStoryPages, model.StoryPage{
						StoryPageId: storyPagesByUuid[data.Fields.Uuid].StoryPageId,
						StoriesId:   storyPagesByUuid[data.Fields.Uuid].StoriesId,
						Text:        data.Fields.Text,
						PageOrder:   data.Fields.Order,
						CreatedAt:   time.Now(),
						Uuid:        data.Fields.Uuid,
						ImagePath:   storyPagesByUuid[data.Fields.Uuid].ImagePath,
					})
				}
			}
			if storiesDbByTitle[key].StartTime != storiesAirtableByTitle[key][0].Fields.StartTime ||
				storiesDbByTitle[key].EndTime != storiesAirtableByTitle[key][0].Fields.EndTime ||
				(storiesAirtableByTitle[key][0].Fields.Image != nil && !strings.Contains(storiesDbByTitle[key].IconPath, storiesAirtableByTitle[key][0].Fields.Image[0].FileName)) {
				updateStories = append(updateStories, storiesDbByTitle[key])
			}
			continue
		}
		story := model.Stories{
			CreatedAt: time.Now(),
			StartTime: storiesAirtableByTitle[key][0].Fields.StartTime,
			EndTime:   storiesAirtableByTitle[key][0].Fields.EndTime,
			Title:     storiesAirtableByTitle[key][0].Fields.Title,
		}
		if storiesAirtableByTitle[key][0].Fields.Icon != nil {
			images := *storiesAirtableByTitle[key][0].Fields.Icon
			file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_STORIES), images[0].FileName, images[0].Url)
			if err != nil {
				log.Println(ctx, "some err while create image", "err", err, "story name", story.Title)
			}
			story.IconPath = file
		}
		createStories = append(createStories, story)
	}

	createStoryPages := make([]model.StoryPage, 0)
	if len(createStories) > 0 {
		createStories, err = h.stories.CreateMany(ctx, createStories)
		if err != nil {
			return err
		}
		for _, story := range createStories {
			storyId := story.StoriesId
			for _, data := range storiesAirtableByTitle[story.Title] {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_STORIES), data.Fields.Image[0].FileName, data.Fields.Image[0].Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "stories name", story.Title)
				}
				createStoryPages = append(createStoryPages, model.StoryPage{
					StoriesId: storyId,
					ImagePath: file,
					Text:      data.Fields.Text,
					PageOrder: data.Fields.Order,
					Uuid:      data.Fields.Uuid,
				})
			}
		}

		err = h.storyPage.CreateMany(ctx, createStoryPages)
		if err != nil {
			return err
		}
	}

	if len(updateStoryPages) > 0 {
		for i := range updateStoryPages {
			if !strings.Contains(updateStoryPages[i].ImagePath, imagesByUuidAirtable[updateStoryPages[i].Uuid][0].FileName) {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_STORIES), imagesByUuidAirtable[updateStoryPages[i].Uuid][0].FileName, imagesByUuidAirtable[updateStoryPages[i].Uuid][0].Url)
				if err != nil {
					log.Println(ctx, "some err while update image", "err", err, "story page uuid", updateStoryPages[i].Uuid)
				}
				updateStoryPages[i].ImagePath = file
			}
		}
		_, err = h.storyPage.UpdateMany(ctx, updateStoryPages)
		if err != nil {
			log.Println(ctx, "error while updating existing posts from airtable:", err)
			return err
		}
	}
	if len(updateStories) > 0 {
		_, err = h.stories.UpdateMany(ctx, updateStories)
		if err != nil {
			return err
		}
	}

	deleteStoryTitles := make([]string, 0)
	for key := range storiesDbByTitle {
		if _, exists := storiesAirtableByTitle[key]; !exists {
			deleteStoryTitles = append(deleteStoryTitles, key)
		}
	}
	if len(deleteStoryTitles) > 0 {
		err = h.stories.DeleteManyByTitle(ctx, deleteStoryTitles)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *AirTableSync) compareHashtags(dbHashtags, airtableHashtags []string) bool {
	if len(dbHashtags) != len(airtableHashtags) {
		return false
	}

	hashtagSet := make(map[string]bool)
	for _, hashtag := range dbHashtags {
		hashtagSet[hashtag] = true
	}

	for _, hashtag := range airtableHashtags {
		if !hashtagSet[hashtag] {
			return false
		}
	}
	return true
}

func (h *AirTableSync) syncHashtags(ctx context.Context) error {
	hashtags, err := h.airTable.GetHashtags(ctx)
	if err != nil {
		return err
	}

	hashtagsAirtableByName := make(map[string]airtable.BaseObject[airtable.Hashtag])
	for _, post := range hashtags {
		hashtagsAirtableByName[post.Fields.Name] = post
	}

	hashtagsDb, err := h.hashtag.GetAll(ctx)
	if err != nil {
		return err
	}

	hashtagsDbByName := make(map[string]model.Hashtag)
	for _, hashtag := range hashtagsDb {
		hashtagsDbByName[hashtag.Name] = hashtag
	}

	createHashtags := make([]model.Hashtag, 0)
	updateHashtags := make([]model.Hashtag, 0)
	deleteHashtags := make([]model.Hashtag, 0)
	for key := range hashtagsAirtableByName {
		if data, ok := hashtagsDbByName[key]; ok {
			var images []airtable.Image
			if hashtagsAirtableByName[key].Fields.Image != nil {
				images = *hashtagsAirtableByName[key].Fields.Image
			}

			if data.NameRu != hashtagsAirtableByName[key].Fields.NameRu ||
				data.NameKz != hashtagsAirtableByName[key].Fields.NameKz ||
				(data.ImagePath == nil && images != nil && len(images) > 0) ||
				(data.ImagePath != nil && images != nil && len(images) > 0 && strings.Contains(*data.ImagePath, images[0].FileName)) {
				data.NameRu = hashtagsAirtableByName[key].Fields.NameRu
				data.NameKz = hashtagsAirtableByName[key].Fields.NameKz
				if (data.ImagePath == nil && images != nil && len(images) > 0) ||
					(data.ImagePath != nil && images != nil && len(images) > 0 && strings.Contains(*data.ImagePath, images[0].FileName)) {
					file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_HASHTAG), images[0].FileName, images[0].Url)
					if err != nil {
						log.Println(ctx, "some err while create image", "err", err, "hashtag name", data.Name)
					}
					data.ImagePath = &file
				}
				updateHashtags = append(updateHashtags, data)
			}
			continue
		}
		hashtag := model.Hashtag{
			NameKz: hashtagsAirtableByName[key].Fields.NameKz,
			NameRu: hashtagsAirtableByName[key].Fields.NameRu,
			Name:   hashtagsAirtableByName[key].Fields.Name,
		}
		if hashtagsAirtableByName[key].Fields.Image != nil && len(*hashtagsAirtableByName[key].Fields.Image) > 0 {
			images := *hashtagsAirtableByName[key].Fields.Image
			file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_HASHTAG), images[0].FileName, images[0].Url)
			if err != nil {
				log.Println(ctx, "some err while create image", "err", err, "hashtag name", hashtag.Name)
			}
			hashtag.ImagePath = &file
		}
		createHashtags = append(createHashtags, hashtag)
	}
	for key := range hashtagsDbByName {
		if _, ok := hashtagsAirtableByName[key]; ok {
			continue
		}
		deleteHashtags = append(deleteHashtags, hashtagsDbByName[key])
	}

	if len(createHashtags) > 0 {
		_, err = h.hashtag.CreateMany(ctx, createHashtags)
		if err != nil {
			return err
		}
	}
	if len(updateHashtags) > 0 {
		_, err = h.hashtag.UpdateMany(ctx, updateHashtags)
		if err != nil {
			return err
		}
	}
	if len(deleteHashtags) > 0 {
		deleteHashtagIds := make([]uint, len(deleteHashtags))
		for i, pr := range deleteHashtags {
			deleteHashtagIds[i] = pr.HashtagID
		}
		err = h.hashtag.DeleteMany(ctx, deleteHashtagIds)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *AirTableSync) syncCollections(ctx context.Context) error {
	collections, err := h.airTable.GetCollections(ctx)
	if err != nil {
		return err
	}

	collectionsAirtableByName := make(map[string]airtable.BaseObject[airtable.Collection])
	for _, post := range collections {
		collectionsAirtableByName[post.Fields.Name] = post
	}

	collectionsDb, err := h.collection.GetAll(ctx)
	if err != nil {
		return err
	}

	collectionDbByName := make(map[string]model.Collection)
	for _, hashtag := range collectionsDb {
		collectionDbByName[hashtag.Name] = hashtag
	}

	createCollections := make([]model.Collection, 0)
	updateCollections := make([]model.Collection, 0)
	deleteCollections := make([]model.Collection, 0)
	for key := range collectionsAirtableByName {
		if data, ok := collectionDbByName[key]; ok {
			var images []airtable.Image
			if collectionsAirtableByName[key].Fields.Image != nil {
				images = *collectionsAirtableByName[key].Fields.Image
			}
			var imagesRu []airtable.Image
			if collectionsAirtableByName[key].Fields.ImageRu != nil {
				imagesRu = *collectionsAirtableByName[key].Fields.ImageRu
			}
			var imagesKz []airtable.Image
			if collectionsAirtableByName[key].Fields.ImageKz != nil {
				imagesKz = *collectionsAirtableByName[key].Fields.ImageKz
			}
			isUpdate := false

			if (data.ImagePath == nil && images != nil && len(images) > 0) ||
				(data.ImagePath != nil && images != nil && len(images) > 0 && strings.Contains(*data.ImagePath, images[0].FileName)) {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), images[0].FileName, images[0].Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "hashtag name", data.Name)
				}
				data.ImagePath = &file
				isUpdate = true
			}
			if (data.ImagePathKz == nil && imagesKz != nil && len(imagesKz) > 0) ||
				(data.ImagePathKz != nil && imagesKz != nil && len(imagesKz) > 0 && strings.Contains(*data.ImagePathKz, imagesKz[0].FileName)) {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), "_kz_"+imagesKz[0].FileName, imagesKz[0].Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "hashtag name", data.Name)
				}
				data.ImagePathKz = &file
				isUpdate = true
			}
			if (data.ImagePathRu == nil && imagesRu != nil && len(imagesRu) > 0) ||
				(data.ImagePathRu != nil && imagesRu != nil && len(imagesRu) > 0 && strings.Contains(*data.ImagePathRu, imagesRu[0].FileName)) {
				file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), "_ru_"+imagesRu[0].FileName, imagesRu[0].Url)
				if err != nil {
					log.Println(ctx, "some err while create image", "err", err, "hashtag name", data.Name)
				}
				data.ImagePathRu = &file
				isUpdate = true
			}
			if data.IsRecommendation != collectionsAirtableByName[key].Fields.IsRecommendation {
				data.IsRecommendation = collectionsAirtableByName[key].Fields.IsRecommendation
				isUpdate = true
			}

			if isUpdate {
				updateCollections = append(updateCollections, data)
			}
			continue
		}
		collection := model.Collection{
			Name:             collectionsAirtableByName[key].Fields.Name,
			IsRecommendation: collectionsAirtableByName[key].Fields.IsRecommendation,
		}
		if collectionsAirtableByName[key].Fields.Image != nil && len(*collectionsAirtableByName[key].Fields.Image) > 0 {
			images := *collectionsAirtableByName[key].Fields.Image
			file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), images[0].FileName, images[0].Url)
			if err != nil {
				log.Println(ctx, "some err while create image", "err", err, "collection name", collection.Name)
			}
			collection.ImagePath = &file
		}
		if collectionsAirtableByName[key].Fields.ImageRu != nil && len(*collectionsAirtableByName[key].Fields.ImageRu) > 0 {
			images := *collectionsAirtableByName[key].Fields.ImageRu
			file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), "_ru_"+images[0].FileName, images[0].Url)
			if err != nil {
				log.Println(ctx, "some err while create image", "err", err, "collection name", collection.Name)
			}
			collection.ImagePath = &file
		}
		if collectionsAirtableByName[key].Fields.ImageKz != nil && len(*collectionsAirtableByName[key].Fields.ImageKz) > 0 {
			images := *collectionsAirtableByName[key].Fields.ImageKz
			file, err := h.storage.CreateImage(ctx, string(model.BUCKET_NAME_COLLECTION), "_kz_"+images[0].FileName, images[0].Url)
			if err != nil {
				log.Println(ctx, "some err while create image", "err", err, "collection name", collection.Name)
			}
			collection.ImagePath = &file
		}
		createCollections = append(createCollections, collection)
	}
	for key := range collectionDbByName {
		if _, ok := collectionsAirtableByName[key]; ok {
			continue
		}
		deleteCollections = append(deleteCollections, collectionDbByName[key])
	}

	if len(createCollections) > 0 {
		_, err = h.collection.CreateMany(ctx, createCollections)
		if err != nil {
			return err
		}
	}
	if len(updateCollections) > 0 {
		_, err = h.collection.UpdateMany(ctx, updateCollections)
		if err != nil {
			return err
		}
	}
	if len(deleteCollections) > 0 {
		deleteCollectionIds := make([]uint, len(deleteCollections))
		for i, pr := range deleteCollections {
			deleteCollectionIds[i] = pr.CollectionID
		}
		err = h.collection.DeleteMany(ctx, deleteCollectionIds)
		if err != nil {
			return err
		}
	}

	return nil
}
