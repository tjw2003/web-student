<script lang="ts">
	import { afterNavigate, goto } from '$app/navigation';
  import axios from 'axios';


  let courses: any = [];
  import { Button } from 'flowbite-svelte';


  // let username = ""

  onMount(() => {
    getselectCourse()
    // username = localStorage.getString("username")
  })


  function getselectCourse() {
    // let username = localStorage.getString('username')
    axios
    .get(
        "http://127.0.0.1:9090/api/course/getselectCourse",
        // JSON.stringify({
        //   username,
        // }),
      {
        headers: {
          Authorization: "Bearer " + localStorage.getItem("prj-jwt"),
        }
      }

      )
      .then((res) => {
        console.log("[then]: ", res);
        courses = res.data.data;
      })
      .catch((err) => {
        console.log("[error]: ", err);
      });
  }

  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from "flowbite-svelte";
  import { onMount } from 'svelte';
</script>

<div class="bg-white border-t p-2">
  <Button on:click={getselectCourse} gradient color="green">刷新</Button>
</div>

<div>
  <Table>
    <TableHead>
      <TableHeadCell>Name</TableHeadCell>
      <TableHeadCell>Credit</TableHeadCell>
      <TableHeadCell>Teacher</TableHeadCell>
    </TableHead>
    <TableBody>
      {#each courses as course}
        <TableBodyRow>
          <TableBodyCell>{course.name}</TableBodyCell>
          <TableBodyCell>{course.credit}</TableBodyCell>
          <TableBodyCell>{course.teacher}</TableBodyCell>
        </TableBodyRow>
      {/each}
    </TableBody>
  </Table>
</div>

<style>
    h1 {
      color: black;
      font-family: "Comic Sans MS", cursive;
      font-size: 2em;
    }
  </style>