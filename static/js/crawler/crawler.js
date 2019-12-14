window.onload = function() {
    Vue.component('crawler-template', {
        props: ['item'],
        template: '#crawler-template'
    })

    var app = new Vue({
        el: '#app',
        data: {
            pages: 1,
            isCrawlerTable: false,
            crawlerData: [
                {}
            ]
        },
        methods: {
            callCrawler() {
                this.isCrawlerTable = false;
                var self = this;
                $.ajax({
                    type: 'POST',
                    url: '/crawler',
                    data: { 'page': this.pages },
                    success: function(data) {
                        self.crawlerData = data;
                        self.isCrawlerTable = true;
                    }
                });
            }
        }
    })
}