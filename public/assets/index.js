(() => {
  // app/assets/index.js
  window.applyFilter = function applyFilter(filterValue) {
    const url = new URL(window.location.href);
    if (url.searchParams.has("filter", filterValue)) {
      url.searchParams.delete("filter", filterValue);
    } else if (!url.searchParams.has("filter")) {
      url.searchParams.set("filter", filterValue);
    } else {
      url.searchParams.append("filter", filterValue);
    }
    window.history.pushState(null, null, url);
    location.reload();
  };
})();
