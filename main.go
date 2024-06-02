package main

import (
	"main/pages/main_page"
	"net/http"
)

func main() {
	init_pages()

	http.ListenAndServe("0.0.0.0:80", nil)
}

func init_pages() {
	main_func_template := main_page.Template()
	http.Handle("/", main_page.Main_page(main_func_template))
}
