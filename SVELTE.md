# SVELTE前端框架

#### 首先使用scoop 下载pnpm(类似npm，效果更好)

```
scoop install nodejs-lts pnpm//会自动下载nodejs环境，如不行
```

nodejs环境scoop下载

```
scoop install nodejs
```

#### 创建svelte框架

```
pnpm create svelte@latest my-app//my-app为项目名称，可替换
```

会有一堆选项，大概根据英语自己选吧

Next steps:

```html
 1: cd web-fronted
 2: npm install (or pnpm install, etc)//安装package.json的依赖
 3: git init && git add -A && git commit -m "Initial commit" (optional)
 4: npm run dev -- --open//运行package.json scripts的内容
```

h 查看帮助

r 启动服务器即可打开网站

cd fronted

pnpm build//生成build文件（具体很复杂，以后再说）

cd..

go run ./

## Svelte组件

首先，向组件添加 `<script>` 标签并声明一个 `name` 变量

```html
<script>
	let name = 'world';
</script>

<h1>Hello world!</h1>
```

可以写成

```
<h1>Hello {name}!</h1>
```

就像可以使用花括号控制文本一样，也可以使用花括号控制元素属性。

```html
<script>
	let src = 'tutorial/image.gif';
</script>

<img src={src} alt="A man dances.">
```

`alt` 属性，该属性用于描述这个标签的用途，以便让使用屏幕阅读器的用户，或网络连接缓慢、不稳定的用户了解这幅图像的用途

```
<img {src} alt="A man dances.">//名称和值相同的属性也可以这么写
```

```
<!-- 注释 -->
```

通常，组件状态的某些部分需要通过 *其它* 部分的计算而得出

```html
<script> 中
let count = 0;
$: doubled = count * 2;

<html> 中
<p>{count} doubled is {doubled}</p> <!--简写-->{count * 2}
```

由于 Svelte 的反应性是由赋值语句触发的，因此使用数组的诸如 `push` 和 `splice` 之类的方法就不会触发自动更新。例如，点击按钮就不会执行任何操作。

解决该问题的一种方法是添加一个多余的赋值语句

```html
function addNumber() {
	numbers.push(numbers.length + 1);
	numbers = numbers;
}
```

但还有一个更惯用的解决方案：

```html
function addNumber() {
	numbers = [...numbers, numbers.length + 1];
}
```

你可以使用类似的模式来替换 `pop`、`shift`、`unshift` 和 `splice` 方法。

In any real application, you'll need to pass data from one component down to its children. To do that, we need to declare *properties*, generally shortened to 'props'. In Svelte, we do that with the `export` keyword. Edit the `Nested.svelte` component:

```html
<script>
	export let answer;
</script>
```

如果你的组件含有有一个对象属性，可以利用`...`语法将它们*spread*（传播）到一个组件上，而不用逐一指定：

```
<Info {...pkg}/>
```

如果你遇到需要遍历的数据列表，请使用 `each` 块：

```html
<ul>
	{#each cats as cat}
		<li><a target="_blank" href="https://www.youtube.com/watch?v={cat.id}">
			{cat.name}
		</a></li>
	{/each}
</ul>
```



表达式 `cats`是一个数组，遇到数组或类似于数组的对象 (即具有`length` 属性)。你都可以通过 `each [...iterable]`遍历迭代该对象。

```html
{#each cats as cat, i}
	<li><a target="_blank" href="https://www.youtube.com/watch?v={cat.id}">
		{i + 1}: {cat.name}
	</a></li>
{/each}
```

##### 事件处理：

```
<div on:mousemove={handleMousemove}>
	The mouse position is {m.x} x {m.y}
</div>
```

div 元素 (或 HTML 文档分区元素) 是一个通用型的流内容容器，在不使用CSS的情况下，其对内容或布局没有任何影响。

带 `once` 修饰符的事件处理程序只运用一次

```
<script>
	function handleClick() {
		alert('no more alerts')
	}
</script>

<button on:click|once={handleClick}>
	Click me
</button>
```

