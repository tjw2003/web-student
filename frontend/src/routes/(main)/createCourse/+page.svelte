<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import axios from "axios";
  import { Button, Card, Input, Label } from "flowbite-svelte";
  import { postByJson } from "../../../lib/axios";

  let message = "";
  let coursename = "";
  let teacher = "";
  let credit = "";

  function create() {
    postByJson("/api/course/createCourse", {
      coursename,
      credit,
      teacher,
    })
      .then((res) => {
        console.log("[then]: ", res);
        alert("你好，成功创建");
      })
      .catch((err) => {
        console.log("[error]: ", err);
        console.log("error info: ", err.response.data.error);
        message = "创建失败";
      });
  }
</script>

<Card>
  <span class="text-2xl text-center font-semibold">创建课程</span>
  <div class="m-2 flex flex-col gap-2">
    <div>
      <Label for="username" class="mb-2">课程名</Label>
      <Input
        type="text"
        id="username"
        placeholder="请输入课程名..."
        required
        bind:value={coursename}
      />
    </div>
    <div>
      <Label for="credit" class="mb-2">课程学分</Label>
      <Input
        type="text"
        id="credit"
        placeholder="请输入课程学分..."
        required
        bind:value={credit}
      />
    </div>
    <div>
      <Label for="teacher" class="mb-2">课程教师</Label>
      <Input
        type="text"
        id="teacher"
        placeholder="请输入课程教师..."
        required
        bind:value={teacher}
      />
    </div>
    <span class="text-red-400">{message}</span>
  </div>

  <div class="flex gap-4 justify-center">
    <Button on:click={create}>创建</Button>
  </div>
</Card>
