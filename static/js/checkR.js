window.onload = function() {

    //query 가져오기
    function getParameterByName(name) {
     name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
     var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
         results = regex.exec(location.search);
     return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
   };
   var cla = getParameterByName("class");
   var row = getParameterByName("row");
   var col = getParameterByName("col");
   document.getElementById(cla+"-"+row+"-"+col).className = "btn_m";
   document.getElementById("reservation").innerText = cla+"-"+String.fromCharCode(row-224) + "-" + (col < 10 ? "0" + col : col);
};
   