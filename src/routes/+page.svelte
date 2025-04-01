<script lang="ts">
    import { echarts } from "$lib/echarts";
    const option = {
        xAxis: {
            type: "category",
            data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        },
        yAxis: {
            type: "value",
        },
        series: [
            {
                data: [150, 230, 224, 218, 135, 147, 260],
                type: "line",
            },
        ],
    };
</script>

<h1>Tea Data Visualization</h1>
<div class="container" use:echarts={option}></div>

<slot />

<style>
    .container {
        width: 500px;
        height: 500px;
    }
</style>
