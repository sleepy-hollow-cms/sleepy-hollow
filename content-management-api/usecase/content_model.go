package usecase

import (
	"context"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/domain"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/port"
	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/log"
)

type ContentModel struct {
	ContentModelPort port.ContentModel
	EntryPort        port.Entry
}

func NewContentModel(
	contentModelPort port.ContentModel,
	entryPort port.Entry,
) *ContentModel {
	return &ContentModel{
		ContentModelPort: contentModelPort,
		EntryPort:        entryPort,
	}
}

func (c *ContentModel) FindByID(id domain.ContentModelID) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.FindByID(context.TODO(), id)

	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	return result, nil
}

func (c *ContentModel) DeleteContentModelByID(id domain.ContentModelID) error {
	contextTodo := context.TODO()

	find, err := c.EntryPort.Find(contextTodo)

	referEntries := find.Filter(func(entry domain.Entry) bool {
		return entry.ContentModelID == id
	})

	if len(referEntries) != 0 {
		return ReferenceByEntryError{
			Reason: "This ContentModel is referenced by Entry and cannot be deleted",
		}
	}

	if err != nil {
		return err
	}

	err = c.ContentModelPort.DeleteByID(contextTodo, id)

	if err != nil {
		log.Logger.Warn(err.Error())
		return err
	}

	return nil
}

func (c *ContentModel) FindContentModelBySpaceID(spaceID domain.SpaceID) (domain.ContentModels, error) {
	models, err := c.ContentModelPort.FindBySpaceID(context.TODO(), spaceID)
	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModels{}, err
	}
	return models, nil
}

func (c *ContentModel) Create(contentModel domain.ContentModel) (domain.ContentModel, error) {
	result, err := c.ContentModelPort.Create(context.TODO(), contentModel)

	if err != nil {
		return domain.ContentModel{}, NewContentModelCreateFailedError(err.Error())
	}

	return result, err
}

func (c *ContentModel) Update(contentModel domain.ContentModel) (domain.ContentModel, error) {
	found, err := c.ContentModelPort.FindByID(context.TODO(), contentModel.ID)
	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	updated := domain.ContentModel{
		ID:        contentModel.ID,
		Name:      contentModel.Name,
		CreatedAt: found.CreatedAt,
		UpdatedAt: contentModel.UpdatedAt,
		Fields:    contentModel.Fields,
	}

	result, err := c.ContentModelPort.Update(context.TODO(), updated)
	if err != nil {
		log.Logger.Warn(err.Error())
		return domain.ContentModel{}, err
	}

	return result, nil
}
