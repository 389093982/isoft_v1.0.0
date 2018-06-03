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


    $(document).ModalEffects({"clearUIFunc":function () {
        // 置空异常信息
        $("#_editEnvInfo_error").html("");
        // 清空表单数据
        clearFormData();
    }});
});

// 清空表单数据
function clearFormData() {
    $("input[id='env_name']").val("");
    $("input[id='env_ip']").val("");
    $("input[id='env_account']").val("");
    $("input[id='env_passwd']").val("");
}

function editEnvInfo() {
    var env_name = $("input[id='env_name']").val();
    var env_ip = $("input[id='env_ip']").val();
    var env_account = $("input[id='env_account']").val();
    var env_passwd = $("input[id='env_passwd']").val();

    $.ajax({
        url:"/env/edit",
        type:"post",
        data:{"env_name":env_name, "env_ip":env_ip, "env_account": env_account, "env_passwd": env_passwd},
        success:function (data) {
            if(data.status=="SUCCESS"){
                window.location.reload();
            }else{
                $("#_editEnvInfo_error").html("*" + data.errorMsg);
            }
        }
    });
}

function connect_test(currentNode) {
    $(currentNode).parents("tr").find(".connect").addClass("loading");
}