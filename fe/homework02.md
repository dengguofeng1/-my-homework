#  个人博客系统开发考核
## :dart: 考核目标
1. 掌握 Next.js 核心功能（SSG/ISR/API Routes）
2. 实现 Supabase 数据集成
3. 完成基础的 CRUD 功能
4. 应用现代化前端开发实践
## :pencil: 任务要求
```bash
# 基础功能（60分）
- [ ] 文章列表页（含分页）
- [ ] 文章详情页 
- [ ] 文章创建页面
- [ ] 基础 SEO 优化
- [ ] 实现 ISR 增量静态再生

# 进阶功能（40分）
- [ ] 文章编辑/删除功能
- [ ] Markdown 内容渲染
- [ ] 标签分类系统
- [ ] 评论功能集成
- [ ] 部署到 Vercel
 ```

## :key: 核心实现要点
### 1. 数据层集成 (15分)
```javascript
// 补充 Supabase 查询示例
export async function getArticles() {
  const { data, error } = await supabase
    .from('articles')
    .select('id, title, slug, created_at')
    .order('created_at', { ascending: false })
  
  if (error) throw new Error('Failed to fetch articles')
  return data
}

// 创建文章示例
export async function createArticle(articleData) {
  const { data, error } = await supabase
    .from('articles')
    .insert([{
      ...articleData,
      author_id: supabase.auth.user()?.id
    }])
  
  if (error) throw new Error('Failed to create article')
  return data
}
 ```


### 2. 列表页优化 (10分)
```javascript
// 添加加载状态和错误处理
function ArticleList() {
  const { data: articles, error } = useSWR('/api/articles', fetcher, {
    refreshInterval: 30000 // 30秒刷新
  })

  if (error) return <ErrorComponent />
  if (!articles) return <LoadingSkeleton />

  return (
    <div className="grid gap-6">
      {articles.map(article => (
        <ArticlePreview key={article.id} article={article} />
      ))}
    </div>
  )
}
 ```


## :bar_chart: 评分标准 考核项 评分标准 功能完整性

基础功能完整度（50%权重） 代码规范

遵循 Next.js 最佳实践（20%权重） 性能优化

Lighthouse 评分 ≥ 90（15%权重） 额外功能

每个实现的进阶功能 +5分
## :tools: 开发路线图
```mermaid
graph TD
    A[项目初始化] --> B[数据库设计]
    B --> C[API 实现]
    C --> D[页面开发]
    D --> E[状态管理]
    E --> F[性能优化]
    F --> G[部署发布]
 ```

## :books: 学习资源
1. Next.js 官方文档
2. Supabase JavaScript 指南
3. SWR 数据请求策略