package usecase

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"admin-pro/internal/domain/entity"
	"admin-pro/internal/domain/repository"
)

type GenUsecase interface {
	GetTableList(ctx context.Context, tableName string, pageSize, pageNo int) ([]*entity.TableInfo, int64, error)
	GenCode(ctx context.Context, tableNames []string) ([]byte, error)
}

type genUsecase struct {
	genRepo repository.GenRepository
}

func NewGenUsecase(genRepo repository.GenRepository) GenUsecase {
	return &genUsecase{
		genRepo: genRepo,
	}
}

func (u *genUsecase) GetTableList(ctx context.Context, tableName string, pageSize, pageNo int) ([]*entity.TableInfo, int64, error) {
	return u.genRepo.GetTableList(ctx, tableName, pageSize, pageNo)
}

func (u *genUsecase) GenCode(ctx context.Context, tableNames []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	defer writer.Close()

	for _, tableName := range tableNames {
		// Get columns
		cols, err := u.genRepo.GetTableColumns(ctx, tableName)
		if err != nil {
			log.Printf("Failed to get columns for table %s: %v", tableName, err)
			continue // skip error table
		}

		// Simple template generation
		// 1. Entity
		entityCode := generateEntity(tableName, cols)
		f, err := writer.Create(fmt.Sprintf("%s/entity.go", tableName))
		if err != nil {
			return nil, fmt.Errorf("create zip entry for %s: %w", tableName, err)
		}

		if _, err := f.Write([]byte(entityCode)); err != nil {
			return nil, fmt.Errorf("write to zip for %s: %w", tableName, err)
		}

		// 2. Repo
		// TODO: Add more files
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close zip writer: %w", err)
	}

	return buf.Bytes(), nil
}

func generateEntity(tableName string, cols []*entity.ColumnInfo) string {
	var sb strings.Builder
	sb.WriteString("package entity\n\n")
	sb.WriteString("type " + toCamelCase(tableName) + " struct {\n")
	for _, col := range cols {
		fieldName := toCamelCase(col.ColumnName)
		fieldType := "string" // simplify
		if strings.Contains(col.DataType, "int") {
			fieldType = "int"
		} else if strings.Contains(col.DataType, "time") {
			fieldType = "time.Time"
		}

		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" gorm:\"column:%s\"` // %s\n",
			fieldName, fieldType, toLowerCamel(fieldName), col.ColumnName, col.ColumnComment))
	}
	sb.WriteString("}\n")
	return sb.String()
}

func toCamelCase(s string) string {
	// Simple impl
	parts := strings.Split(s, "_")
	var res string
	for _, p := range parts {
		if len(p) > 0 {
			res += strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return res
}

func toLowerCamel(s string) string {
	c := toCamelCase(s)
	if len(c) > 0 {
		return strings.ToLower(c[:1]) + c[1:]
	}
	return c
}
