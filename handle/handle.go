package handle

import (
    "net/http"
    "fmt"
    "log"
    "bytes"
)

type HTML_PAGE_t struct {
    Title string
}

var page HTML_PAGE_t

func InitFrontPage( title string ) {
    page.Title = title
}

func checkRequestToken( r *http.Request  ) bool {
    cookie, err := r.Cookie( "token" )
    if err != nil {
        log.Println(err)
        return false
    }

    log.Printf("token cookie:%+v", cookie )
    return true
}


func AdminHandler(w http.ResponseWriter, r *http.Request) {
    valid_token := checkRequestToken( r )
    if ! valid_token {
        // to login
        var b bytes.Buffer
        err := t_login.Execute( &b, page )
        if err != nil {
            fmt.Fprintf( w , err.Error() )
            return
        }
        b.WriteTo(w)
        return
    }
    fmt.Fprintf(w, "admin" )
}
