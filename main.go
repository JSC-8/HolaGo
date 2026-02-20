package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	//cargar la api key de .env
	error := godotenv.Load()
	if error != nil {
		log.Fatal("Error al cargar .env")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")

	if apiKey == "" {
		log.Fatal("Variable API KEY IS NULL")
	}
	//Interactuar con la API mandando la Key
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	//Esta parte es para listar las versiones que tenemos disponibles con nuestra version
	/*fmt.Println("Buscando modelos disponibles...")
	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("- ", m.Name)
	} //Fin mostrar versiones gemini
	*/

	model := client.GenerativeModel("models/gemini-3-flash-preview")

	//Escribir un promp para gemini

	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Escribe algo para guglu")
		input, _ := reader.ReadString('\n')
		fmt.Println("Analizando tu consulta...")

		resp, err := model.GenerateContent(ctx, genai.Text(input))
		if err != nil {
			log.Fatal(err)
		}

		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					fmt.Println(part)
				}

			}
		}
	}

}
