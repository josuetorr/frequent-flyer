document.body.addEventListener("htmx:afterRequest", (e) => {
  const elSelector = e.detail.xhr.getResponseHeader("HX-FOCUS");
  if (!elSelector) return;

  const el = document.querySelector(elSelector);
  el.focus();
});
