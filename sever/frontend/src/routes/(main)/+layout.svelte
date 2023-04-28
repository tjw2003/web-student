<script>
  import { goto } from "$app/navigation";
  import { page } from "$app/stores";

  import { Button } from "flowbite-svelte";
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger,
    Dropdown,
    DropdownItem,
    Chevron,
    DropdownDivider,
  } from "flowbite-svelte";

//   $: activeUrl = ;
//   $: console.log(activeUrl)
  $: active1 = ($page.url.pathname == '/studentpage')
</script>

<Navbar>
  <NavBrand href="/">
    <span
      class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
    >
      TJW选课系统{$page.url.pathname}
    </span>
  </NavBrand>
  <NavHamburger />
  <NavUl>
    <NavLi href="/studentpage"
      >所有课程</NavLi
    >
    <NavLi href="/createCourse"
      >创建课程</NavLi
    >
    <NavLi href="/searchCourse"
      >查询已选课程</NavLi
    >
    <NavLi
      class="hover:cursor-pointer"
      on:click={() => {
        localStorage.removeItem("prj-jwt");
        localStorage.removeItem("prj-id");
        goto("/");
      }}>退出</NavLi
    >
  </NavUl>
</Navbar>

<div class="flex flex-col items-stretch">
  <slot />
</div>
