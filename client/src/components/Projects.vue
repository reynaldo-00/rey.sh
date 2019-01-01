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
