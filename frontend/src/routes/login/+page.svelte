<!--js文件-->
<script lang="ts">
  import { Input, Label, Card, Button } from "flowbite-svelte";

  let username = "";
  let password = "";

  import axios from "axios";
  import { goto } from "$app/navigation";

  function onLoggedIn(data: any) {
    console.log("[onLoggedIn]: ", data);
    localStorage.setItem("prj-jwt", data.token);
    localStorage.setItem("prj-id", data.user.id);
    localStorage.setItem("username", data.user.username);
    goto("/studentpage");
  }

  function login() {
    axios
      .post(
        "http://127.0.0.1:9090/api/user/login",
        JSON.stringify({
          username,
          password,
        })
      )
      .then((res) => {
        console.log("[then]: ", res);
        res = res.data;
        onLoggedIn(res.data);
      })
      .catch((err) => {
        console.log("[error]: ", err);
        console.log("error info: ", err.response.data.error);
        Err();
      });
  }

  function Err() {
    message = "用户名不存在或密码错误";
  }

  let message = "";
</script>

<Card
  class="absolute left-1/2 top-1/2 translate-x-[-50%] translate-y-[-50%] w-2/5 flex flex-col gap-2"
>
  <span class="text-2xl text-center font-semibold">学生选课系统</span>
  <div class="m-2 flex flex-col gap-2">
    <div class="mb-3">
      <Label for="username" class="mb-2">账号</Label>
      <Input
        type="text"
        id="username"
        placeholder="请输入账号..."
        required
        bind:value={username}
      />
    </div>
    <div>
      <Label for="password" class="mb-2">密码</Label>
      <Input
        type="password"
        id="password"
        placeholder="请输入密码..."
        required
        bind:value={password}
      />
    </div>
    <span class="text-red-400">{message}</span>
  </div>

  <div class="flex gap-4 justify-center">
    <Button on:click={login}>登录</Button>
    <Button color="alternative" on:click={() => goto("/signup")}>注册</Button>
  </div>
</Card>
