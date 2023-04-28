<script>
  import { goto } from "$app/navigation";

  let username = "";
  let password = "";
  let passwordAgain = "";
  let message = "";

  import axios from "axios";
  import { Button, Card, Input, Label } from "flowbite-svelte";

  function Err (){
    message = "两次输入密码不同"
  }

  function usersignout(){
    alert("你好，成功注册");
  }

  function Errow (){
    message = "用户名已存在"
  }
  function signout() {
    if ((password == passwordAgain)) {
      axios
        .post(
          "http://127.0.0.1:9090/api/user/register",
          JSON.stringify({
            username,
            password,
          })
        )
        .then((res) => {
          console.log(res);
          usersignout()
        })
        .catch((err) => {
          console.log(err);
          Errow()
        });
    }else{
      Err()
	}
  }
</script>

<Card>
  <span class="text-2xl text-center font-semibold">学生选课系统</span>
  <div class="m-2 flex flex-col gap-2">
    <div class="mb-3">
      <Label for="username" class="mb-2">账号</Label>
      <Input type="text" id="username" placeholder="请输入账号..." required bind:value={username}/>
    </div>
    <div class="mb-3">
      <Label for="password" class="mb-2">密码</Label>
      <Input type="password" id="password" placeholder="请输入密码..." required bind:value={password} />
    </div>
    <div>
      <Label for="password_again" class="mb-2">确认密码</Label>
      <Input type="password" id="password_again" placeholder="请再次输入密码..." required bind:value={passwordAgain} />
    </div>
    <span class="text-red-400">{message}</span>
  </div>

  <div class="flex gap-4 justify-center">
    <Button color="alternative" on:click={() => goto('/')}> 登录 </Button>
    <Button
      on:click={() => {
        console.log(username);
        console.log(password);
        signout();
      }}
    >
      注册
    </Button>

    <!-- <div class="text-center ...">
      <button
        class="bg-red-500 hover:bg-red-700 ..."
        on:click={() => {
          // TODO: login
          console.log(username);
          console.log(password);
          signout();
          // login()
        }}
        id="btn"
      >
        注册
      </button>
      <a href="../">
        <button class="bg-red-500 hover:bg-red-700 ..."> 登录 </button>
      </a> -->
  </div>
</Card>

<div class="m-auto">
  <div class="text-center ... m-5 ... flex flex-col gap-2">
    <input
      bind:value={username}
      class="text-center ... border-2 border-red-500 ..."
      style="display: block"
      type="text"
      id="input-username"
      name="username"
      placeholder="请输入账号..."
    />
    <input
      bind:value={password}
      class="text-center ... border-2 border-red-500 ..."
      style="display: block"
      type="password"
      id="input-password"
      name="password"
      placeholder="请输入密码..."
    />
    <input
      bind:value={passwordAgain}
      class="text-center ... border-2 border-red-500 ..."
      type="password"
      id="input-password"
      name="password"
      placeholder="请再次输入密码..."
    />
    <span class="text-red-400">{message}</span>
  </div>

  <!-- <input type="button" value="登录"/> -->
  <!-- svelte-ignore missing-declaration -->
  <div class="text-center ...">
    <button
      class="bg-red-500 hover:bg-red-700 ..."
      on:click={() => {
        // TODO: login
        console.log(username);
        console.log(password);
        signout();
        // login()
      }}
      id="btn"
    >
      注册
    </button>
    <a href="../">
      <button class="bg-red-500 hover:bg-red-700 ..."> 登录 </button>
    </a>
  </div>
</div>
