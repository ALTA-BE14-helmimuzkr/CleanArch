package handler

import "api/features/book"

type AddUpdateBookRequest struct {
	Judul       string `json:"judul"`
	TahunTerbit int    `json:"tahun_terbit"`
	Penulis     string `json:"penulis"`
}

func ToCore(data interface{}) *book.Core {
	res := book.Core{}

	switch data.(type) {
	case AddUpdateBookRequest:
		cnv := data.(AddUpdateBookRequest)
		res.Judul = cnv.Judul
		res.TahunTerbit = cnv.TahunTerbit
		res.Penulis = cnv.Penulis
	default:
		return nil
	}

	return &res
}
