<!--
  ~ Licensed to the Apache Software Foundation (ASF) under one or more
  ~ contributor license agreements.  See the NOTICE file distributed with
  ~ this work for additional information regarding copyright ownership.
  ~ The ASF licenses this file to You under the Apache License, Version 2.0
  ~ (the "License"); you may not use this file except in compliance with
  ~ the License.  You may obtain a copy of the License at
  ~
  ~     http://www.apache.org/licenses/LICENSE-2.0
  ~
  ~ Unless required by applicable law or agreed to in writing, software
  ~ distributed under the License is distributed on an "AS IS" BASIS,
  ~ WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  ~ See the License for the specific language governing permissions and
  ~ limitations under the License.
-->
<template>
  <div class="__container_services_index">
    <search-table :search-domain="searchDomain">
      <template #bodyCell="{ column, record, text }">
        <template v-if="column.dataIndex === 'serviceName'">
          {{ record.versionGroup }}
          <span class="service-link" @click="viewDistribution(text, text.versionGroupValue)">
            <b>
              <Icon style="margin-bottom: -2px" icon="material-symbols:attach-file-rounded"></Icon>
              {{ text }}
            </b>
          </span>
        </template>

        <template v-else-if="column.dataIndex === 'versionGroupSelect'">
          <a-select :value="text?.versionGroupValue" :bordered="false" style="width: 80%">
            <a-select-option
              v-for="(item, index) in text?.versionGroupArr"
              :value="item"
              :key="index"
            >
              {{ item }}
            </a-select-option>
          </a-select>
        </template>
      </template>
    </search-table>
  </div>
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { nextTick, provide, reactive, watch } from 'vue'
import { searchService } from '@/api/service/service'
import { SearchDomain } from '@/utils/SearchUtil'
import SearchTable from '@/components/SearchTable.vue'
import { PROVIDE_INJECT_KEY } from '@/base/enums/ProvideInject'
import { PRIMARY_COLOR } from '@/base/constants'
import { Icon } from '@iconify/vue'
import { queryMetrics } from '@/base/http/promQuery'
import { promQueryList } from '@/utils/PromQueryUtil'

let __null = PRIMARY_COLOR
const router = useRouter()
const route = useRoute()
let query = route.query['query']
const columns = [
  {
    title: 'service',
    key: 'service',
    dataIndex: 'serviceName',
    sorter: true,
    width: '30%',
    ellipsis: true
  },
  {
    title: 'versionGroup',
    key: 'versionGroup',
    dataIndex: 'versionGroupSelect',
    width: '25%'
  },
  {
    title: 'avgQPS',
    key: 'avgQPS',
    dataIndex: 'avgQPS',
    sorter: true,
    width: '15%'
  },
  {
    title: 'avgRT',
    key: 'avgRT',
    dataIndex: 'avgRT',
    sorter: true,
    width: '15%'
  },
  {
    title: 'requestTotal',
    key: 'requestTotal',
    dataIndex: 'requestTotal',
    sorter: true,
    width: '15%'
  }
]

// Extract version and group.
// Todo
// const extractVersionAndGroup = (input:string) =>{
//   console.log(input)
//   return
//   const regex = /version:\s*(\d+\.\d+),\s*group:\s*(\d+\.\d+)/;
//   const match = input.match(regex);
//   console.log(match)
//   if (match) {
//     return {
//       version: match[1],
//       group: match[2]
//     };
//   } else {
//     return null;
//   }
// }

const handleResult = (result: any) => {
  return result.map((service: any) => {
    service.versionGroupSelect = {}
    service.versionGroupSelect.versionGroupArr = service.versionGroups.map((item: any) => {
      return (item.versionGroup =
        (item.version ? 'version: ' + item.version + ', ' : '') +
          (item.group ? 'group: ' + item.group : '') || '无')
    })
    service.versionGroupSelect.versionGroupValue = service.versionGroupSelect.versionGroupArr[0]
    return service
  })
}

function serviceInfo(params: any, table: any) {
  return searchService(params).then(async (res) => {
    return promQueryList(res, ['avgQPS', 'avgRT', 'requestTotal'], async (service: any) => {
      service.avgQPS = await queryMetrics(
        `sum (dubbo_provider_qps_total{interface='${service.serviceName}'}) by (interface)`
      )
      service.avgRT = await queryMetrics(
        `avg(dubbo_consumer_rt_avg_milliseconds_aggregate{interface="${service.serviceName}",method=~"$method"}>0)`
      )
      service.requestTotal = await queryMetrics(
        `sum (increase(dubbo_provider_requests_total{interface="${service.serviceName}"}[1m]))`
      )
    })
  })
}

const searchDomain = reactive(
  new SearchDomain(
    [
      {
        label: 'serviceName',
        param: 'keywords',
        placeholder: 'typeAppName',
        defaultValue: query,
        style: {
          width: '200px'
        }
      }
    ],
    serviceInfo,
    columns,
    undefined,
    undefined,
    handleResult
  )
)

searchDomain.onSearch(handleResult)

const viewDetail = (serviceName: string) => {
  router.push({ name: 'detail', params: { pathId: serviceName } })
}

const viewDistribution = (serviceName: string, versionAndGroup: string) => {
  // let group = extractVersionAndGroup(versionAndGroup)?.group || ''
  // let version = extractVersionAndGroup(versionAndGroup)?.version || ''
  router.push({ path: `/resources/services/distribution/${serviceName}` })
}

provide(PROVIDE_INJECT_KEY.SEARCH_DOMAIN, searchDomain)
watch(route, (a, b) => {
  searchDomain.queryForm['keywords'] = a.query['query']
  searchDomain.onSearch()
  console.log(a)
})
</script>
<style lang="less" scoped>
.__container_services_index {
  .service-link {
    padding: 4px 10px 4px 4px;
    border-radius: 4px;
    color: v-bind('PRIMARY_COLOR');

    &:hover {
      cursor: pointer;
      background: rgba(133, 131, 131, 0.13);
    }
  }
}
</style>
