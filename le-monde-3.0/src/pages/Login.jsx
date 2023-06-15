import React, { useEffect }  from 'react';
import Logo from '../assets/LeMonde3-0.png'

function Login() {
    useEffect(() => {
    }, []);
  
    return (
      <div class="flex h-screen w-full items-center justify-center bg-gray-900 bg-cover bg-no-repeat" style={{backgroundImage :"url('https://images.unsplash.com/photo-1504711434969-e33886168f5c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1170&q=80')"}}>
        <div class="rounded-xl bg-gray-800 bg-opacity-50 px-16 py-10 shadow-lg backdrop-blur-md max-sm:px-8">
          <div class="text-white">
            <div class="mb-8 flex flex-col items-center">
              <img src={Logo} width="350" alt="" srcset="" />
            </div>
            <form action="#">
              <div class="flex mb-4 text-lg justify-center">
                <input class="rounded-3xl border-none bg-sky-800 bg-opacity-50 px-6 py-2 text-center text-inherit placeholder-slate-500 shadow-lg outline-none backdrop-blur-md" type="text" placeholder="email" />
              </div>

              <div class="flex mb-4 text-lg justify-center">
                <input class="rounded-3xl border-none bg-sky-800 bg-opacity-50 px-6 py-2 text-center text-inherit placeholder-slate-500 shadow-lg outline-none backdrop-blur-md" type="Password" placeholder="password" />
              </div>
              <div class="mt-8 flex justify-center text-lg text-black">
                <button type="submit" class="rounded-3xl bg-black bg-opacity-50 px-10 py-2 text-white shadow-xl backdrop-blur-md transition-colors duration-300 hover:bg-yellow-600">Login</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    );
}

export default Login;