$(function () {
    // 定义一个全局的 vueData,初始数据为空
    var envInfosVueData = {
        envInfos:[]
    };
    // 定义一个全局的 vue 实例,引用这个全局的 vueData
    var envInfosVue = new Vue({
        // 修改 vue 默认分隔符,解决冲突问题
        delimiters: ['[[', ']]'],
        el: '#envInfo_list',
        data: envInfosVueData
    });

    function pageToolFunction(obj) {
        // 渲染分页信息
        $('#pageTool').Paging({pagesize: obj.paginator.pagesize,count:obj.paginator.totalcount,current:1,callback:function(page,size,count){
                loadPageData(page, size, null);
        }});
    }

    function loadPageData(current_page,page_size,pageToolFunction) {
        $.ajax({
            url:"/env/list",
            method:"post",
            data:{"current_page":current_page, "page_size":page_size},
            async: false,
            success:function (data) {
                var obj = JSON.parse(data);
                // 使用 $set 去修改这个 vueData 进行刷新页面
                envInfosVue.$set(envInfosVueData, 'envInfos', obj.envInfos);
                if(pageToolFunction != null){
                    pageToolFunction(obj);              // 渲染分页
                }
            }
        });
    }
    // 加载第一页,10条记录,加载完成之后使用 pageToolFunction 函数进行分页渲染
    loadPageData(1,10,pageToolFunction);
});