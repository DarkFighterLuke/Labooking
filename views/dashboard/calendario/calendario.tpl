<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div>
    <div id="calendar"></div>

    <script type="text/javascript" src="/js/calendario/components/underscore/underscore-min.js"></script>
    <script type="text/javascript" src="/js/calendario/js/calendar.js"></script>
    <script type="text/javascript">
        var calendar = $("#calendar").calendar({
            tmpl_path: "/js/calendario/tmpls/",
            view: "month",
            events_source: [
                {{range .Eventi}
        }
        {
            "id"
        :
            {
                {.
                    Id
                }
            }
        ,
            "title"
        :
            {
                {.
                    Title
                }
            }
        ,
            "url"
        :
            {
                {.
                    Url
                }
            }
        ,
            "class"
        :
            {
                {.
                    Class
                }
            }
        ,
            "start"
        :
            {
                {.
                    Start
                }
            }
        ,
            "end"
        :
            {
                {.
                    End
                }
            }
        }
        ,
        {
            {
                end
            }
        }
        ]
        })
        ;
    </script>
</div>