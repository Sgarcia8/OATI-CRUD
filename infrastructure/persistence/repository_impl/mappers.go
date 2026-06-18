package repository_impl

import (
	"oati-crud-comentarios/domain"
	"oati-crud-comentarios/infrastructure/persistence/models_orm"
)

func toCommentORM(c *domain.Comment) *models_orm.CommentORM {
	return &models_orm.CommentORM{
		Id:        c.Id,
		Content:   c.Content,
		Tutorial:  &models_orm.TutorialORM{Id: c.TutorialId},
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		IsDeleted: c.IsDeleted,
	}
}

func toCommentDomain(c *models_orm.CommentORM) *domain.Comment {
	tutorialId := 0
	if c.Tutorial != nil {
		tutorialId = c.Tutorial.Id
	}
	return &domain.Comment{
		Id:         c.Id,
		Content:    c.Content,
		TutorialId: tutorialId,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		IsDeleted:  c.IsDeleted,
	}
}

func toCommentDomainList(list []*models_orm.CommentORM) []*domain.Comment {
	result := make([]*domain.Comment, len(list))
	for i, c := range list {
		result[i] = toCommentDomain(c)
	}
	return result
}

func toTutorialORM(t *domain.Tutorial) *models_orm.TutorialORM {
	return &models_orm.TutorialORM{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		PublishedAt: t.PublishedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		IsDeleted:   t.IsDeleted,
	}
}

func toTutorialDomain(t *models_orm.TutorialORM) *domain.Tutorial {
	tutorial := &domain.Tutorial{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		PublishedAt: t.PublishedAt,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		IsDeleted:   t.IsDeleted,
	}
	if len(t.Comments) > 0 {
		tutorial.Comments = toCommentDomainList(t.Comments)
		for _, c := range tutorial.Comments {
			if c.TutorialId == 0 {
				c.TutorialId = t.Id
			}
		}
	}
	return tutorial
}

func toTutorialDomainList(list []*models_orm.TutorialORM) []*domain.Tutorial {
	result := make([]*domain.Tutorial, len(list))
	for i, t := range list {
		result[i] = toTutorialDomain(t)
	}
	return result
}
