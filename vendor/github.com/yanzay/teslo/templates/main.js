window.addEventListener("load", function(evt) {
  var ws = new WebSocket("ws://localhost:8080/ws");
  ws.onopen = function(e) {
    console.log("OPEN");
  };
  ws.onclose = function(e) {
    console.log("CLOSE");
  };
  ws.onmessage = function(e) {
    var message = JSON.parse(e.data);
    var el = document.getElementById(message.id)
    el.outerHTML = message.content;
  };
  ws.onerror = function(e) {
    console.log("Error: ", e.data);
  };
  var app = document.getElementById("app");

  var parentIds = function(el) {
    return $(el).parents().map(function(i, el) {return el.id;}).toArray().filter(function(id) {return id !== ""});
  };

  var clickHandler = function(e) {
    console.log("Click handler");
    console.log(e.target.id);
    console.log(parentIds);
    if (e.target.id) {
      ws.send(JSON.stringify({event: "click", id: e.target.id, parents: parentIds(e.target)}));
    }
  };
  var submitHandler = function(e) {
    console.log("Submit handler");
    e.preventDefault();
    var form = $(e.target).serialize();
    console.log($(e.target));
    console.log(form);
    var resp = {
      event: "submit",
      id: e.target.id,
      data: form,
    };
    console.log(resp);
    ws.send(JSON.stringify(resp));
  };
  var changeHandler = function(e) {
    console.log("Change handler");
    console.log(e.target);
    console.log($(e.target));
    console.log($(e.target).data());
    console.log(e.target.checked);
    if (e.target.checked !== undefined) {
      var data = $(e.target).data();
      data.checked = e.target.checked;
      var resp = {
        event: "change",
        parents: parentIds(e.target),
        data: JSON.stringify($(e.target).data())
      };
      ws.send(JSON.stringify(resp));
    }
  };
  app.addEventListener("click", clickHandler);
  app.addEventListener("submit", submitHandler);
  app.addEventListener("change", changeHandler);
});

