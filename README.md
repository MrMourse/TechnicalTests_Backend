# TechnicalTests_Backend

Technical tests from Scalingo, main goal is to provide a utility: browser.

This application was made with Golang idea.

It uses 4 files :

    3 Html files:

        -head : Print the beginning of the web page, it concerns title, head and introduce the boostrap framework
        -data : Print the data gathered by the search method on main.go
        -foot : Print the footer of the page

    1 Go file:

        It's the main application, we can separate it in 3 parts:

            -The Get request : it just a print of head and foot, the main goal is to catch the query

            -The Search request:

                -it obtains the query
                -it gets the 100 last repository
                -it gathers the hash_table language and url for each repository
                -it gathers all the information of the repositories
                -it prints the html page result (data.html)

            -The server part:

                with the ServeAndListen method

To use the application please be on localhost:3000/