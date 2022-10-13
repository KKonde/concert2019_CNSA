window.onload = function() {

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



    document.getElementById('card').style.zIndex = "2";
    document.getElementById('open').style.opacity = 0;
    document.getElementById('logo').style.opacity = 1;
    document.getElementById('logo2').style.opacity = 1;
    document.getElementById('front').style.transform = "rotateY(180deg)";
    document.getElementById('back').style.transform = "rotateY(360deg)";
    setTimeout(function(){


      document.getElementById('password-textbox2').classList.add('on2');
      document.getElementById('show').classList.add('button1');
      document.getElementById('show2').classList.add('button2');
      document.getElementById('password-textbox3').classList.add('on2');
      document.getElementById('change-button').classList.add('on2');
      document.getElementById('p').classList.add('on2');

    },1500);


    document.getElementById('change-button').onclick = function(){
      a = document.getElementById('password-textbox2').value;
      b = document.getElementById('password-textbox3').value;
      if((a == "") ||( b == "")){
        document.getElementById('warning').style.opacity=1;
        document.getElementById('warning').textContent="비밀번호를 입력해주세요!";
        document.getElementById('warning').style.color= "red";
        document.getElementById('change-button').classList.remove('on2');
        document.getElementById('change-button').style.opacity=1;
        document.getElementById('change-button').classList.remove('vibration');
        void document.getElementById('change-button').offsetWidth;
        document.getElementById('change-button').classList.add('vibration');
      }

      else if(a != b){
            document.getElementById('warning').style.opacity=1;
              document.getElementById('warning').textContent="비밀번호가 일치하지 않습니다!";
                document.getElementById('warning').style.color= "red";
            document.getElementById('change-button').classList.remove('on2');
            document.getElementById('change-button').style.opacity=1;
            document.getElementById('change-button').classList.remove('vibration');
            void document.getElementById('change-button').offsetWidth;
            document.getElementById('change-button').classList.add('vibration');
          }

          else {
                        document.getElementById('warning').textContent="비밀번호가 일치합니다!";
                        document.getElementById('warning').style.color= "green";
                      }
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

  };
