function TocviewToggle(glyph, id) {
  var s = document.getElementById(id).style;
  var expand = s.display == "none";
  s.display = expand ? "block" : "none";
  glyph.innerHTML = expand ? "&#9660;" : "&#9658;";
}
