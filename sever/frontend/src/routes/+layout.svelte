<script>
  import { afterNavigate, goto } from "$app/navigation";
    import "../app.css";

	afterNavigate((nav) => {
		// console.log(nav);
		// console.log(localStorage.getItem('prj-jwt'));
		if (localStorage.getItem('prj-jwt') == null) {
			if (nav.to?.route.id != '/' && nav.to?.route.id != '/signout') goto('/');
		}
	});
  import { DarkMode } from 'flowbite-svelte';
</script>

<div class="absolute inset-0 bg-tjw-sb -z-50" />
<div class="absolute inset-0 bg-white/40 -z-20 dark:bg-black" />

<div class="h-full w-full flex flex-col items-strech">
    <slot />
	<DarkMode />
</div>
