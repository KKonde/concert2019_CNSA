window.onload = function(){
 //query 가져오기
 function getParameterByName(name) {
  name = name.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]");
  var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
      results = regex.exec(location.search);
  return results === null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
};
var a = getParameterByName("error");
if(a=="password"){
  document.getElementById('warning').style.opacity=1;
  document.getElementById('warning').textContent="현재 비밀번호가 일치하지 않습니다!";
  document.getElementById('warning').style.color= "red";
}
else if (a=="blank"){
  document.getElementById('warning').style.opacity=1;
  document.getElementById('warning').textContent="새 비밀번호를 입력해 주세요!";
  document.getElementById('warning').style.color= "red";

} else if (a=="check"){
  document.getElementById('warning').style.opacity=1;
  document.getElementById('warning').textContent="새 비밀번호가 일치하지 않습니다!";
  document.getElementById('warning').style.color= "red";

}

    
     
     
      





  document.getElementById('show').onclick = function(){


    document.getElementById('password-textbox2').classList.toggle('active');
    if(document.getElementById('password-textbox2').classList.contains('active')){
       document.getElementById('password-textbox2').type = "text";
       document.getElementById('show').style.backgroundColor = "#d6d6d6";
     }
    else {
      document.getElementById('password-textbox2').type = "password";
      document.getElementById('show').style.backgroundColor = "#525050";
    }
  }

  document.getElementById('show2').onclick = function(){


    document.getElementById('password-textbox3').classList.toggle('active');
    if(document.getElementById('password-textbox3').classList.contains('active')){
       document.getElementById('password-textbox3').type = "text";
       document.getElementById('show2').style.backgroundColor = "#d6d6d6";
     }
    else {
      document.getElementById('password-textbox3').type = "password";
      document.getElementById('show2').style.backgroundColor = "#525050";
    }
  }
  document.getElementById('show3').onclick = function(){


    document.getElementById('password-textbox1').classList.toggle('active');
    if(document.getElementById('password-textbox1').classList.contains('active')){
       document.getElementById('password-textbox1').type = "text";
       document.getElementById('show3').style.backgroundColor = "#d6d6d6";
     }
    else {
      document.getElementById('password-textbox1').type = "password";
      document.getElementById('show3').style.backgroundColor = "#525050";
    }
  }

};
