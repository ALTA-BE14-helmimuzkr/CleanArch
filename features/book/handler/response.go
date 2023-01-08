package handler

import "api/features/book"

type BookResponse struct {
	ID          uint   `json:"id"`
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
	Pemilik     string `json:"pemilik"`
}
type AddBookResponse struct {
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}

func ToResponse(feature string, book book.Core) interface{} {
	switch feature {
	case "add":
		return AddBookResponse{
			Judul:       book.Judul,
			TahunTerbit: book.TahunTerbit,
			Penulis:     book.Penulis,
		}
	default:
		return BookResponse{
			ID:          book.ID,
			Judul:       book.Judul,
			TahunTerbit: book.TahunTerbit,
			Penulis:     book.Penulis,
			Pemilik:     book.Pemilik,
		}
	}
}

func ToListResponse(dataCore []book.Core) []BookResponse {
	var ResponData []BookResponse

	for _, value := range dataCore {
		ResponData = append(ResponData, ToResponse("list", value).(BookResponse))
	}
	return ResponData
}
