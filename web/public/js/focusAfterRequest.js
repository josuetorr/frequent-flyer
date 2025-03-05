document.body.addEventListener("htmx:afterRequest", (e) => {
  const elId = e.detail.xhr.getResponseHeader("HX-FOCUS");
  if (!elId) return;

  const el = document.getElementById(elId);
  el.focus();
});
