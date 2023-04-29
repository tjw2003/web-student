<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import axios from "axios";
  import { Button, Card, Input, Label } from "flowbite-svelte";
  // afterNavigate((nav) => {
  //   console.log(nav);
  //   console.log(localStorage.getItem("prj-jwt"));
  //   if (localStorage.getItem("prj-jwt") == null) {
  //     if (nav.to?.route.id != "/login" && nav.to?.route.id != "/register")
  //       goto("../");
  //   }
  // });

  function createcourse() {
    alert("你好，成功创建");
  }
  function error() {
    message = "创建失败";
  }
  let message = "";
  let coursename = "";
  let teacher = "";
  let credit = "";

  function create() {
    axios
      .post(
        "http://127.0.0.1:9090/api/course/createCourse",
        JSON.stringify({
          coursename,
          credit,
          teacher,
        }),
        {
          headers: {
            Authorization: "Bearer " + localStorage.getItem("prj-jwt"),
          },
        }
      )
      .then((res) => {
        console.log("[then]: ", res);
        createcourse();
      })
      .catch((err) => {
        console.log("[error]: ", err);
        console.log("error info: ", err.response.data.error);
        error();
      });
  }
</script>

<Card
>
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

<!-- <div class="m-auto">
  <div class="text-center ... m-5 flex flex-col gap-2">
    <label for="exampleInputEmail1">课程名</label>
    <input
      bind:value={coursename}
      class="text-center ... border-2 border-blue-500 ..."
      style="display: block"
      type="text"
      id="input-coursename"
      name="coursename"
      placeholder="请输入课程名..."
    />
  </div>

  <div class="text-center ... m-5 flex flex-col gap-2">
    <label for="exampleInputEmail1">课程学分</label>
    <input
      bind:value={credit}
      class="text-center ... border-2 border-blue-500 ..."
      style="display: block"
      type="text"
      id="input-credit"
      name="credit"
      placeholder="请输入课程学分..."
    />
  </div>

  <div class="text-center ... m-5 flex flex-col gap-2">
    <label for="exampleInputEmail1">课程教师</label>
    <input
      bind:value={teacher}
      class="text-center ... border-2 border-blue-500 ..."
      style="display: block"
      type="text"
      id="input-teacher"
      name="crteacherdit"
      placeholder="请输入课程教师..."
    />
  </div> -->
  <!-- <span class="text-red-400">{message}</span> -->

  <!-- <div class="text-center ...">
    <button
      on:click={() => {
        create();
      }}
      id="btn"
      class="bg-red-500 hover:bg-red-700 ..."
    >
      创建课程
    </button>
  </div> -->
<!-- </div> -->

<style>
  h1 {
    color: black;
    font-family: "Comic Sans MS", cursive;
    font-size: 2em;
  }
</style>
