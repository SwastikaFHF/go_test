
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <title>在线markdown编辑器</title>
    <meta name="keywords" content="online markdown editor,markdown编辑器,在线编辑markdown" />
    <meta name="description" content="在线编辑markdown的工具" />
    <meta name="COPYRIGHT" content="版权所有2016,http://www.aitangba.com/">
    <link href="./css/mahua.css?v=3" rel="stylesheet" type="text/css" />
    <link rel="icon" href="./icon/go.ico">
</head>
<body>
    <div class="mahua-container">
        <div id="mahua-docname" class="mahua-docanme"></div>
        <ul class="mahua-topbar" id="mahua-topbar">
            <li class="mahua-config" id="mahua-config"></li>
            <li class="mahua-session" id="mahua-session"></li>
        </ul>
        <div class="mahua-toolbar" style="display:none" id="mahua-toolbar">
            <table cellspacing="0">
                <tr>
                    <td> 
                        <label for="mahua-changeSession">选择文档</label>
                    </td>
                    <td>
                        <select id="mahua-changeSession" name=""> </select> 
                        <a href="javascript:;" id="mahua-removeSession">删除</a>
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <table cellspacing="0">
                            <tr>
                                <td>
                                    <label for="mahua-linenum">行号<label>
                                </td>
                                <td>
                                    <input id="mahua-linenum" type="checkbox" checked class="checkbox" value="" name="" />
                                </td>
                            </tr>
                        </table>
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <input id="mahua-importSource" type="button" class="button" value="导入" name="" />
                        <input id="mahua-exportHTML" type="button" class="button" value="导出" name="" />
                        <input id="mahua-closeConfig" type="button" class="button" value="关闭设置" name="" />
                    </td>
                </tr>
            </table>
        </div>
        <div class="mahua-spliter" id="mahua-spliter"></div>
        <div class="mahua-left">
            <div id="mahua-editor" style="height:100%;width:50%;"></div>
        </div>
        <div class="mahua-right" id="mahua-result"> </div>
    </div>
    <script type="text/javascript" src="./js/ace.js"></script>
    <script type="text/javascript" src="./js/keybinding-vim.js"></script>
    <script type="text/javascript" src="./js/all.js?v=5"></script>
    <script type="text/javascript">
        (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'hm.baidu.com/h.js?f628d86243daf05c564aa17f55e27b02'; 
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s); 
        })();
    </script>
</body>
</html>
