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
                type: "scatter",
            },
        ],
    };
</script>
<div class="container" use:echarts={option}></div>