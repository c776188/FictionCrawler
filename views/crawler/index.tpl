<!DOCTYPE html>
<html>

<head>
    <title>Crawler</title>
    <script src="https://code.jquery.com/jquery-3.4.1.js" integrity="sha256-WpOohJOqMqqyKL9FccASB9O0KwACQJpFTUBLTYOVvVU=" crossorigin="anonymous"></script>
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
    <script src="/static/js/crawler/crawler.js"></script>

    <style type="text/css">
        [v-cloak] {
            display: none;
        }
    </style>
</head>

<body>
    <div id="app">
        <br/> 撈幾頁: <input type="text" id="pages" v-model="pages">
        <input type="button" value="Search" @click="callCrawler ">
        <br/>
        <table v-cloak v-if="isCrawlerTable" cellspacing="0" cellpadding="4" border="1">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                </tr>
            </thead>
            <tbody id="rows">
                <tr is="crawler-template" v-for="(item, index) in crawlerData" :item="item" :key="index"></tr>
            </tbody>
        </table>
    </div>
</body>

<script type="text/x-template" id="crawler-template">
    <tr>
        <td>{{item.Id}}</td>
        <td>
            <a :href="'http://big5.quanben.io' + item.Url" target="_blank"> {{item.Name}} </a>
        </td>
    </tr>
</script>

</html>