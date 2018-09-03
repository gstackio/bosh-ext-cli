package web2

const logsViewerJS string = `
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/ace.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/theme-monokai.js" type="text/javascript" charset="utf-8"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/ace/1.3.3/mode-javascript.js" type="text/javascript" charset="utf-8"></script>
<script type="text/javascript">


function getUrlParameter(sParam) {
    var sPageURL = decodeURIComponent(window.location.search.substring(1)),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : sParameterName[1];
        }
    }
};

$(function() {
    var taskID = getUrlParameter("task-id");
    var requestArguments = [
      {"name":"id","value":taskID},
      {"name":"debug"}
    ];

	var depsReq = {"command":"task","arguments":requestArguments}
    $.post("/api/command", JSON.stringify(depsReq))
    .done(function(data) {
    	var debugLogs = data.Blocks.join();
        
    	var editor = ace.edit("editor");
		var session = editor.session
		session.insert({
		   row: session.getLength(),
		   column: 0
		}, "\n" + debugLogs)

	    editor.setTheme("ace/theme/terminal");
	    editor.session.setMode("ace/mode/less");
	    editor.setFontSize(11);
	    editor.setReadOnly(true);
    });
});

</script>



`
