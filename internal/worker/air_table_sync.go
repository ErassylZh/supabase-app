package worker

import (
	"context"
	"log"
	"work-project/internal/airtable"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type AirTableSync struct {
	airTable    repository.AirTable
	product     repository.Product
	post        repository.Post
	storage     repository.StorageClient
	image       repository.Image
	hashtag     repository.Hashtag
	postHashtag repository.PostHashtag
}

func NewAirTableSync(
	airTable repository.AirTable,
	product repository.Product,
	post repository.Post,
	storage repository.StorageClient,
	image repository.Image,
	hashtag repository.Hashtag,
	postHashtag repository.PostHashtag) *AirTableSync {
	return &AirTableSync{
		airTable:    airTable,
		product:     product,
		post:        post,
		storage:     storage,
		image:       image,
		hashtag:     hashtag,
		postHashtag: postHashtag,
	}
}

func (h *AirTableSync) Run() (err error) {
	ctx := context.Background()
	if err := h.syncProducts(ctx); err != nil {
		log.Println("error while syncing products:", err)
		return err
	}

	if err := h.syncPosts(ctx); err != nil {
		log.Println("error while syncing posts:", err)
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
			if post.Company != postsAirtableByUuid[uuid].Fields.Company ||
				post.Language != postsAirtableByUuid[uuid].Fields.Language ||
				post.Title != postsAirtableByUuid[uuid].Fields.Title ||
				post.Description != postsAirtableByUuid[uuid].Fields.Description ||
				post.Status != postsAirtableByUuid[uuid].Fields.Status ||
				post.Body != postsAirtableByUuid[uuid].Fields.Body ||
				post.ReadTime != postsAirtableByUuid[uuid].Fields.ReadTime ||
				post.Point != postsAirtableByUuid[uuid].Fields.Point ||
				post.QuizTime != postsAirtableByUuid[uuid].Fields.QuizTime ||
				post.RatingStatus != postsAirtableByUuid[uuid].Fields.RatingStatus {

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
				updatePosts = append(updatePosts, post)
			}
			if !h.compareHashtags(existsHashtags, postsAirtableByUuid[uuid].Fields.HashtagName) {
				err = h.postHashtag.DeleteByPostId(ctx, post.PostID)
				if err != nil {
					return err
				}
				var postHashtags []model.PostHashtag
				for _, name := range postsAirtableByUuid[uuid].Fields.HashtagName {
					ht, err := h.hashtag.GetByName(ctx, name)
					if err != nil {
						return err
					}
					postHashtags = append(postHashtags, model.PostHashtag{
						PostId:    post.PostID,
						HashtagId: ht.HashtagID,
					})
				}
				_, err = h.postHashtag.CreateMany(ctx, postHashtags)
				if err != nil {
					return err
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
		_, err = h.postHashtag.CreateMany(ctx, postHashtags)
		if err != nil {
			log.Println(ctx, "error while create post hashtags from airtable ", "err", err)
			return err
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
