<!DOCTYPE html>
<html>
<head>
    {{.HtmlHead}}
</head>

<body>
{{.Nav}}
{{.Sidebar}}

<div class="col-sm-9 col-sm-offset-3 col-lg-10 col-lg-offset-2 main">
    {{.Header}}
    {{.LayoutContent}}
</div>	<!--/.main-->

{{.Footer}}
</body>

</html>
