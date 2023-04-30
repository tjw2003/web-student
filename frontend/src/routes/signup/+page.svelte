<script>
  import { goto } from "$app/navigation";

  let username = "";
  let password = "";
  let passwordAgain = "";
  let message = "";

  import { postByJson } from "../../lib/axios";
  import { Button, Card, Input, Label } from "flowbite-svelte";

  function signup() {
    if (password == passwordAgain) {
      postByJson("/api/user/register", {
        username,
        password,
      })
        .then((res) => {
          console.log(res);
          alert("你好，成功注册");
        })
        .catch((err) => {
          console.log(err);
          message = "用户名已存在";
        });
    } else {
      message = "两次输入密码不同";
    }
  }
</script>

<Card class="absolute left-1/2 top-1/2 translate-x-[-50%] translate-y-[-50%] w-2/5 flex flex-col gap-2">
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
    <div class="mb-3">
      <Label for="password" class="mb-2">密码</Label>
      <Input
        type="password"
        id="password"
        placeholder="请输入密码..."
        required
        bind:value={password}
      />
    </div>
    <div>
      <Label for="password_again" class="mb-2">确认密码</Label>
      <Input
        type="password"
        id="password_again"
        placeholder="请再次输入密码..."
        required
        bind:value={passwordAgain}
      />
    </div>
    <span class="text-red-400">{message}</span>
  </div>

  <div class="flex gap-4 justify-center">
    <Button color="alternative" on:click={() => goto("/login")}>登录</Button>
    <Button on:click={signup}>注册</Button>
  </div>
</Card>
