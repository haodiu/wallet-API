package services

import (
	"context"
	"database/sql"
	"time"

	"wallet/models"
)

type DocumentServiceImpl struct {
	documentDB *sql.DB
	ctx        context.Context
}

func (d DocumentServiceImpl) CreateDocument(document *models.Document) error {
	_, errExec := d.documentDB.ExecContext(d.ctx, "INSERT INTO document (docType, docNumber, issuingAuthority, expiryDate, img, userid, create_at) VALUES (?, ?, ?, ?, ?, ?, ?)", document.DocType, document.DocNumber, document.IssuingAuthority, document.ExpiryDate, document.Img, document.UserID, time.Now())
	if errExec != nil {
		return errExec
	}
	return nil
}

func (d DocumentServiceImpl) GetDocuments(value string) ([]models.Document, error) {
	var documents []models.Document
	rows, errQuery := d.documentDB.QueryContext(d.ctx, "SELECT * FROM document WHERE id = ? OR issuingAuthority = ?", value, value)
	if errQuery != nil {
		return documents, errQuery
	}
	for rows.Next() {
		var document models.Document
		err := rows.Scan(&document.ID, &document.DocType, &document.DocNumber, &document.IssuingAuthority, &document.ExpiryDate, &document.Img, &document.UserID, &document.CreateAt, &document.UpdateAt)
		if err != nil {
			return documents, err
		}
		documents = append(documents, document)
	}
	return documents, nil
}

func (d DocumentServiceImpl) DeleteDocument(id string) error {
	_, err := d.documentDB.ExecContext(d.ctx, "DELETE FROM document WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func NewDocumentService(documentDB *sql.DB, ctx context.Context) DocumentService {
	return &DocumentServiceImpl{
		documentDB: documentDB,
		ctx: ctx,
	}
}