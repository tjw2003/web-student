<script lang="ts">
  import { afterNavigate, goto } from "$app/navigation";
  import axios from "axios";

  let courses: any = [];

  function selectright() {
    alert("你好，成功选课");
  }

  function getcourse() {
    axios
      .get("http://127.0.0.1:9090/api/course", {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("prj-jwt"),
        },
      })
      .then(function (response) {
        console.log(response.data);
        courses = response.data.data;
      })
      .catch((err) => {
        console.log(err);
      });
  }

  let teacher = "";
  let coursename = "";
  let credit = "";

  function selectCourse() {
    // let username = localStorage.getString("username");
    axios
      .post(
        "http://127.0.0.1:9090/api/course/selectCourse",
        JSON.stringify({
          // username,
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
        selectright();
      })
      .catch((err) => {
        console.log("[error]: ", err);
      });
  }

  import {
    Card,
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from "flowbite-svelte";

  import { Button } from "flowbite-svelte";
  import { onMount } from "svelte";

  onMount(() => {
    getcourse()
  })
</script>

<div class="bg-white border-t p-2">
  <Button on:click={getcourse}>刷新</Button>
</div>
<!-- <div> -->
  <!-- svelte-ignore missing-declaration -->
  <!-- <button
    on:click={getcourse}
    class="bg-red-500 hover:bg-red-700 p-1 rounded shadow"
  >
    查看所有课程
  </button>
</div> -->

<div>
  <Table>
    <TableHead>
      <TableHeadCell>ID</TableHeadCell>
      <TableHeadCell>Name</TableHeadCell>
      <TableHeadCell>Credit</TableHeadCell>
      <TableHeadCell>Teacher</TableHeadCell>
      <TableHeadCell>Button</TableHeadCell>
    </TableHead>
    <TableBody>
      {#each courses as course}
        <TableBodyRow>
          <TableBodyCell>{course.id}</TableBodyCell>
          <TableBodyCell>{course.name}</TableBodyCell>
          <TableBodyCell>{course.credit}</TableBodyCell>
          <TableBodyCell>{course.teacher}</TableBodyCell>
          <TableHeadCell>
            <!-- svelte-ignore missing-declaration -->
            <Button
              on:click={() => {
                coursename = course.name;
                credit = course.credit;
                teacher = course.teacher;
                selectCourse();
              }}
              gradient
              color="green"
              pill={true}
            >
              选课</Button
            >
          </TableHeadCell>
        </TableBodyRow>
      {/each}
    </TableBody>
  </Table>
</div>

<style>
  h1 {
    color: black;
    font-family: "Comic Sans MS", cursive;
    font-size: 4em;
  }
</style>
