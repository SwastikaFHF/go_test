<!DOCTYPE html>
<html lang="en">
<head>
	<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta charset="utf-8">
    <meta name="description" content="JSON Lint is a web based validator and reformatter for JSON, a lightweight data-interchange format.">

    <title>JSONLint PRO - The JSON Validator</title>

    <link rel="stylesheet" href="css/jsl-screen.css" type="text/css" media="screen, projection">
    <link href="css/jsl-jquery-linedtextarea.css" type="text/css" rel="stylesheet">
    <link rel="icon" type="image/png" href="icon/go.ico" />
</head>
    <body>
		<input type="hidden" id="reformat" value="1" />
		<input type="hidden" id="compress" value="0" />

		<div id="json-composite-placeholder">

        <div id="results_header" class="hide">
            <h3>
                Results <img title="Loading..." class="reset" alt="Loading" src="images/loadspinner.gif" id="loadSpinner" name="loadSpinner">
            </h3>
        </div>

        <script type="text/javascript">
            if (typeof JSON === 'undefined') {
                document.write('<sc' + 'ript type="text/javascript" src="js/utils/json2.js"></sc' + 'ript>');
            }
        </script>
        <script src="js/jsl.main.js"></script>
        <script src="js/jsl.format.js"></script>
        <script src="js/jsl.parser.js"></script>
	</body>
</html>