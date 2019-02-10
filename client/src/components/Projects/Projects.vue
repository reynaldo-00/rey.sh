<template>
    <div>
        <ProjectDetails 
            v-for="project in projects"
            :project="project"
            :key="project.id"
        />
    </div>
</template>

<script>
import axios from 'axios';

import ProjectDetails from './ProjectDetails'

export default {
    name: 'Projects',
    components: {
        ProjectDetails
    },
    data: () => ({
        projects: Array
    }),
    methods: {
        getProjects: async function()  {
            const resp = await axios('http://localhost:9001/repos');
            this.projects = resp.data;
        },
    },
    beforeMount() {
        this.getProjects();
    }
}
</script>

<style>

</style>
