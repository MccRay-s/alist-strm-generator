import { http } from './http'

export class ConfigAPI {
  private baseUrl = '/config'

  /**
   * 创建配置
   */
  async create(data: Api.Config.Create) {
    return http.post<Api.Config.Record, Api.Config.Create>(this.baseUrl, data)
  }

  /**
   * 更新配置
   */
  async update(id: number, data: Api.Config.Update) {
    return http.put(`${this.baseUrl}/${id}`, data)
  }

  /**
   * 删除配置
   */
  async delete(id: number) {
    return http.delete(`${this.baseUrl}/${id}`)
  }

  /**
   * 查询所有配置
   */
  async configs() {
    return http.get<Api.Config.Record[]>(`${this.baseUrl}/list`)
  }

  /**
   * 通过配置代码查询配置项
   * @param code 配置项的代码
   */
  async getByCode(code: string) {
    return http.get<Api.Config.Record>(`${this.baseUrl}/code/${code}`)
  }
}

export const configAPI = new ConfigAPI()
