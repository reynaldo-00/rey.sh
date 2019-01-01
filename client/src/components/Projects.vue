<template>
    <div>

    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Projects',
    data: () => ({
        projects: Array
    }),
    methods: {
        getProjects: async function()  {
            const projectsResponse = await axios('http://localhost:9001/repos');
            const pro = projectsResponse.data.viewer.pinnedRepositories.nodes;
            this.projects = pro;
        },
    },
    beforeMount() {
        this.getProjects();
    }
}
</script>

<style>

</style>
