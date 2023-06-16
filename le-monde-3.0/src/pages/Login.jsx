import React, { useEffect }  from 'react';
import Logo from '../assets/LeMonde3-0.png'
import { useNavigate } from "react-router-dom";
import axios from 'axios'

function Login() {
  let navigate = useNavigate(); 

    useEffect(() => {
    }, []);

    async function login() {
      const user_email = document.getElementById("email")
      console.log(user_email.value)

      const apiData = await axios.get("http://localhost:8080/api/admin/database/User",
      {
        headers: {
          Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2ODY5MjEwNDcsInVzZXJfaWQiOjB9.1QD5509nEYHxEW5tmavmQG9tCvnN3vlzd06VOXbfR80'
        },
        params: {
          email: user_email.value
        }
      }).then(function (response) {
        navigate("/Home");
        console.log(response?.status + ": " + JSON.stringify(response?.data));
      }).catch((err) => {
        closeOpenPopUp("Incorrect email");
        document.getElementById("myForm").reset();
        console.log("Error " + err?.response?.status + ": " + err?.response?.data);
      });

      console.log(apiData)

      // if (apiData.response.Status === 404) {
        //    closeOpenPopUp("Unknown Email")
        //    user_email.reset();
      //}
      //console.log(apiData.response)
      //console.log(apiData)
    }

    function closeOpenPopUp(Error_string) {
      const display = document.getElementById('popup-modal').style.display;
      display === "block" ? document.getElementById('popup-modal').style.display='none' : document.getElementById('popup-modal').style.display='block'
      document.getElementById('modal_text').innerHTML = Error_string;
    }

    const navigateRegister = (FileName) => {
      navigate("/ItAPI");
    }
  
    return (
      <div>
        <div id="popup-modal" class="fixed w-100% h-100% bg-black/70 p-4 overflow-x-hidden backdrop-blur-md bg-cover bg-no-repeat overflow-y-auto md:inset-0 h-[calc(100%)] max-h-full z-10" style={{display: "none"}}>
          <div class="fixed top-2/4 left-1/2 -translate-x-1/2 -translate-y-1/2 w-full max-w-md max-h-full">
              <div class="relative items-center right-50% bg-white rounded-xl shadow dark:bg-black/50">
                  <div class="p-6 text-center">
                      <svg aria-hidden="true" class="mx-auto mb-4 text-red-600 w-14 h-14 dark:text-red-500/80" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                      <h3 id="modal_text" class="mb-5 text-2xl font-normal text-TextColor dark:text-gray-100">Error</h3>
                      <button onClick={closeOpenPopUp} type="button" class="text-white bg-amber-600/80 hover:bg-amber-600 focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-amber-600 font-medium rounded-xl text-sm inline-flex items-center px-5 py-2.5 text-center">
                          Modify
                      </button>
                  </div>
              </div>
          </div>
        </div>
        <div class="flex h-screen w-full items-center justify-center bg-gray-900 bg-cover bg-no-repeat" style={{backgroundImage :"url('https://images.unsplash.com/photo-1504711434969-e33886168f5c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80')"}}>
          <div class="rounded-xl bg-gray-800 bg-opacity-50 px-16 pt-10 pb-4 shadow-lg backdrop-blur-md max-sm:px-8">
            <div class="text-white">
              <div class="mb-8 flex flex-col items-center">
                <img src={Logo} width="350" alt="" srcset="" />
              </div>
              <form id="myForm">
                <div class="flex mb-4 text-lg justify-center">
                  <input id="email" class="rounded-2xl border-none bg-sky-800 bg-opacity-50 px-6 py-2 text-center text-inherit placeholder-slate-500 shadow-lg outline-none backdrop-blur-md" type="text" placeholder="email" />
                </div>

                <div class="flex mb-4 text-lg justify-center">
                  <input id="password" class="rounded-2xl border-none bg-sky-800 bg-opacity-50 px-6 py-2 text-center text-inherit placeholder-slate-500 shadow-lg outline-none backdrop-blur-md" type="Password" placeholder="password" />
                </div>
                <div class="mt-8 flex justify-center text-lg text-black">
                  <button onClick={login} class="rounded-2xl bg-black bg-opacity-50 px-10 py-2 text-white shadow-xl backdrop-blur-md transition-colors duration-300 hover:bg-sky-600">Login</button>
                </div>
              </form>
            </div>
            <span class="text-white text-sm flex pt-4 justify-center items-center text-center cursor-pointer">
              Let's create an account
            </span>
          </div>
        </div>
      </div>
    );
}

export default Login;