window.onload = function() {
    Vue.component('crawler-template', {
        props: ['item'],
        template: '#crawler-template'
    })

    var app = new Vue({
        el: '#app',
        data: {
            pages: 1,
            crawlerData: [
                {}
            ]
        },
        methods: {
            callCrawler() {
                var self = this;
                $.ajax({
                    type: 'POST',
                    url: '/crawler',
                    data: { 'page': this.pages },
                    success: function(data) {
                        self.crawlerData = data;
                    }
                });
            }
        }
    })
}