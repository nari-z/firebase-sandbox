package main

import (
	"fmt"
	"os"
  
	"golang.org/x/net/context"

	"google.golang.org/api/iterator"
	"cloud.google.com/go/firestore"
)

func main() {
	projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		fmt.Errorf("Failed initializing Firestore", err)
		return
	}
	defer client.Close()

	// add
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
	})
	if err != nil {
		fmt.Errorf("Failed adding alovelace: %v", err)
		return
	}

	// list
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Errorf("Failed to iterate: %v", err)
			return
		}
		fmt.Println(doc.Data())
	}

	fmt.Println("Done.")
}