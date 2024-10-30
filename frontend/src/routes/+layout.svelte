<script lang="ts">
    import {onMount} from "svelte";
    import { Test } from "$core/teststruct"
    let text : String = "None";

    function fetchProjects() {
        Test().then((res) => {
            text = res;
            console.log(res);
        });
    }

    onMount(() => {
        fetchProjects();
    });
</script>

<h1>Text : {text}</h1>

<slot></slot>

<style>

</style>