package main

import (
	"github.com/aaronland/go-storage"	
	sfom_storage "github.com/sfomuseum/go-sfomuseum-storage"
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dsn := flag.String("dsn", "storage=fs root=.", "...")
	prefix := flag.String("prefix", "", "...")	
	
	flag.Parse()

	var store storage.Store
	var err error

	if *prefix != "" {
		store, err = sfom_storage.NewStoreWithPrefix(*dsn, *prefix)
	} else {
		store, err = sfom_storage.NewStore(*dsn)
	}
	
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()

		fname := filepath.Base(path)
		key := fname

		err = store.Put(key, fh)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("stored %s in %s\n", path, store.URI(key))
	}
}
