## Dev-Client-Go

**dev-client-go** is a client library for the Forem (dev.to) [developer api](https://developers.forem.com/api) written in Go. It provides fully typed methods for every operation you can carry out with the current api (beta)(0.9.7)

### Installation
> Go version >= 1.13
```sh
$ go get github.com/Mayowa-Ojo/dev-client-go
```

### Usage
Import the package and initialize a new client with your auth token(api-key).
To get a token, see the authentication [docs](https://developers.forem.com/api#section/Authentication)
```go
package main

import (
   dev "github.com/Mayowa-Ojo/dev-client-go"
)

func main() {
   token := <your-api-key>
   client, err := dev.NewClient(token)
   if err != nil {
      // handle err
   }
}
```

<hr style="border:1px solid gray"> </hr>

### Documentation
Examples on basic usage for some of the operations you can carry out.

#### Articles [[API doc](https://developers.forem.com/api#tag/articles)]
Articles are all the posts that users create on DEV that typically show up in the feed.

Example:

Get published articles

query parameters gives you options to filter the results 
```go
// ...
// fetch 10 published articles
articles, err := client.GetPublishedArticles(
   dev.ArticleQueryParams{
      PerPage: 10
   }
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Articles: \n%+v", articles)
// ...
```

Create an article

you can pass the article content as string or as a markdown file by passing the `filepath` as a second parameter
```go
// ...
payload := dev.ArticleBodySchema{}
payload.Article.Title = "The crust of structs in Go"
payload.Article.Published = false
payload.Article.Tags = []string{"golang"}

article, err := client.CreateArticle(payload, "article_sample.md")
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Article: \n%+v", article)
// ...
```

#### Organizations [[API doc](https://developers.forem.com/api#tag/organizations)]
Example:

Get an organization
```go
// ...
organization, err := c.GetOrganization(orgname)
if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Article: \n%+v", organization)
// ...
```

Get users in an organization
```go
// ...
users, err := client.GetOrganizationUsers(
   orgname,
   OrganizationQueryParams{
      Page:    1,
      PerPage: 5,
   },
)

if err != nil {
   fmt.Println(err.Error())
}

fmt.Printf("Article: \n%+v", organization)
// ...
```