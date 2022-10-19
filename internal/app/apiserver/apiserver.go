package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/vitalis-virtus/http-rest-api/internal/app/model"
	"github.com/vitalis-virtus/http-rest-api/internal/app/store"
	"io"
	"log"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter()}
}

// Start  -- the function with which we will start the http server and connect ot DB
func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/author/", s.GetAuthors()).Methods("GET")

	s.router.HandleFunc("/author/{id}", s.GetAuthor()).Methods("GET")

	s.router.HandleFunc("/author/", s.CreateAuthor()).Methods("POST")

	s.router.HandleFunc("/author/{id}", s.DeleteAuthor()).Methods("DELETE")

	s.router.HandleFunc("/author/{id}", s.UpdateAuthor()).Methods("PUT")

}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

// GetAuthors handle function
func (s *APIServer) GetAuthors() http.HandlerFunc {
	// here we can initialize variables only for this handler
	var authors []model.Author

	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.store.DB.Query("SELECT * FROM booksdb.authors")
		if err != nil {
			log.Fatal(err)
		}

		defer res.Close()

		for res.Next() {
			var author model.Author
			err := res.Scan(&author.ID, &author.Name, &author.Surname)
			if err != nil {
				log.Fatal(err)
			}
			authors = append(authors, author)
		}

		resp, _ := json.Marshal(authors)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

		err = res.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

}

// GetAuthor handle function
func (s *APIServer) GetAuthor() http.HandlerFunc {
	// here we can initialize variables only for this handler
	var author model.Author

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		res, err := s.store.DB.Query(fmt.Sprintf("SELECT * FROM booksdb.authors WHERE id=%v", authorID))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Close()

		for res.Next() {
			err := res.Scan(&author.ID, &author.Name, &author.Surname)
			if err != nil {
				log.Fatal(err)
			}
		}

		resp, _ := json.Marshal(author)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)

		err = res.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

}

// UpdateAuthor handle function
func (s *APIServer) UpdateAuthor() http.HandlerFunc {
	// here we can initialize variables only for this handler
	var name, surname string

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		name = r.FormValue("name")
		surname = r.FormValue("surname")

		if name != "" {
			updateName, err := s.store.DB.Query(fmt.Sprintf("UPDATE booksdb.authors SET name='%s' WHERE id=%v", name, authorID))
			if err != nil {
				log.Fatal(err)
			}
			defer updateName.Close()
		}

		if surname != "" {
			updateSurname, err := s.store.DB.Query(fmt.Sprintf("UPDATE booksdb.authors SET surname='%s' WHERE id=%v", surname, authorID))
			if err != nil {
				log.Fatal(err)
			}
			defer updateSurname.Close()
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

}

// DeleteAuthor handle function
func (s *APIServer) DeleteAuthor() http.HandlerFunc {
	// here we can initialize variables only for this handler
	var author model.Author

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		authorID := vars["id"]

		res, err := s.store.DB.Query(fmt.Sprintf("SELECT * FROM booksdb.authors WHERE id=%v", authorID))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Close()

		for res.Next() {
			err := res.Scan(&author.ID, &author.Name, &author.Surname)
			if err != nil {
				log.Fatal(err)
			}
		}

		resD, err := s.store.DB.Query(fmt.Sprintf("DELETE FROM booksdb.authors WHERE id=%v", authorID))
		if err != nil {
			log.Fatal(err)
		}

		defer resD.Close()

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = res.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

}

// CreateAuthor handle function
func (s *APIServer) CreateAuthor() http.HandlerFunc {
	// here we can initialize variables only for this handler
	var name, surname string

	return func(w http.ResponseWriter, r *http.Request) {
		name = r.FormValue("name")
		surname = r.FormValue("surname")
		if name == "" || surname == "" {
			fmt.Println("Bad request")

		} else {
			insert, err := s.store.DB.Query(fmt.Sprintf("INSERT INTO booksdb.authors (`name`, `surname`) VALUES ('%s', '%s')", name, surname))

			if err != nil {
				log.Fatal(err)
			}
			defer insert.Close()
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
		}

	}
}

// handleHello handle function
func (s *APIServer) handleHello() http.HandlerFunc {
	type request struct {
		name string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
