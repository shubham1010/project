package controllers

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Hello World using mux of controllers"))
}
