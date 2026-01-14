import requests
import json
import pytest
import random
import string

BASE_URL = "http://localhost:8080"

@pytest.fixture(scope="module")
def api_client():
    """创建测试客户端"""
    client = BlogAPITester()
    yield client

class BlogAPITester:
    def __init__(self):
        self.token = None
        self.user_id = None
        self.post_id = None
        # 生成随机用户名和邮箱避免冲突
        random_suffix = ''.join(random.choices(string.ascii_lowercase + string.digits, k=6))
        self.username = f"testuser_{random_suffix}"
        self.email = f"test_{random_suffix}@example.com"
        self.password = "test123"
        
    def test_welcome(self):
        """测试欢迎接口"""
        response = requests.get(f"{BASE_URL}/")
        assert response.status_code == 200
        data = response.json()
        assert "message" in data
        assert "version" in data
        print(f"响应: {data}")
    
    def test_register(self):
        """测试用户注册"""
        data = {
            "username": self.username,
            "password": self.password,
            "email": self.email
        }
        response = requests.post(f"{BASE_URL}/register", json=data)
        assert response.status_code == 201
        result = response.json()
        assert "message" in result
        print(f"注册成功: {self.username}")
    
    def test_login(self):
        """测试用户登录"""
        data = {
            "username": self.username,
            "password": self.password
        }
        response = requests.post(f"{BASE_URL}/login", json=data)
        assert response.status_code == 200
        result = response.json()
        assert "token" in result
        self.token = result["token"]
        print(f"Token 已保存: {self.token[:20]}...")
    
    def test_create_post(self):
        """测试创建文章"""
        assert self.token is not None, "需要先登录获取 token"
        
        headers = {"Authorization": f"Bearer {self.token}"}
        data = {
            "title": "测试文章",
            "content": "这是一篇测试文章的内容"
        }
        response = requests.post(f"{BASE_URL}/posts", json=data, headers=headers)
        assert response.status_code == 201
        result = response.json()
        assert "ID" in result
        self.post_id = result["ID"]
        print(f"文章 ID: {self.post_id}")
    
    def test_get_posts(self):
        """测试获取所有文章"""
        response = requests.get(f"{BASE_URL}/posts")
        assert response.status_code == 200
        result = response.json()
        assert isinstance(result, list)
        print(f"文章数量: {len(result)}")
    
    def test_get_post(self):
        """测试获取单篇文章"""
        assert self.post_id is not None, "需要先创建文章"
        
        response = requests.get(f"{BASE_URL}/posts/{self.post_id}")
        assert response.status_code == 200
        result = response.json()
        assert result["ID"] == self.post_id
        print(f"获取文章: {result['title']}")
    
    def test_update_post(self):
        """测试更新文章"""
        assert self.post_id is not None and self.token is not None
        
        headers = {"Authorization": f"Bearer {self.token}"}
        data = {
            "title": "更新后的标题",
            "content": "更新后的内容"
        }
        response = requests.put(f"{BASE_URL}/posts/{self.post_id}", json=data, headers=headers)
        assert response.status_code == 200
        result = response.json()
        assert result["title"] == "更新后的标题"
        print(f"更新成功: {result['title']}")
    
    def test_create_comment(self):
        """测试创建评论"""
        assert self.post_id is not None and self.token is not None
        
        headers = {"Authorization": f"Bearer {self.token}"}
        data = {
            "content": "这是一条测试评论",
            "post_id": self.post_id
        }
        response = requests.post(f"{BASE_URL}/comments", json=data, headers=headers)
        assert response.status_code == 201
        result = response.json()
        assert "ID" in result
        print(f"评论 ID: {result['ID']}")
    
    def test_get_comments(self):
        """测试获取文章评论"""
        assert self.post_id is not None and self.token is not None
        
        headers = {"Authorization": f"Bearer {self.token}"}
        response = requests.get(f"{BASE_URL}/comments/post/{self.post_id}", headers=headers)
        assert response.status_code == 200
        result = response.json()
        assert isinstance(result, list)
        assert len(result) > 0
        print(f"评论数量: {len(result)}")
    
    def test_delete_post(self):
        """测试删除文章"""
        assert self.post_id is not None and self.token is not None
        
        headers = {"Authorization": f"Bearer {self.token}"}
        response = requests.delete(f"{BASE_URL}/posts/{self.post_id}", headers=headers)
        assert response.status_code == 200
        result = response.json()
        assert "message" in result
        print(f"删除成功")
    
    def test_unauthorized_access(self):
        """测试未授权访问"""
        data = {"title": "未授权文章", "content": "这应该失败"}
        response = requests.post(f"{BASE_URL}/posts", json=data)
        assert response.status_code == 401
        result = response.json()
        assert "error" in result
        print(f"未授权访问被正确拒绝")
    
# pytest 测试用例（按照执行顺序）
def test_01_welcome(api_client):
    """测试欢迎接口"""
    api_client.test_welcome()

def test_02_register(api_client):
    """测试用户注册"""
    api_client.test_register()

def test_03_login(api_client):
    """测试用户登录"""
    api_client.test_login()

def test_04_unauthorized_access(api_client):
    """测试未授权访问"""
    api_client.test_unauthorized_access()

def test_05_create_post(api_client):
    """测试创建文章"""
    api_client.test_create_post()

def test_06_get_posts(api_client):
    """测试获取所有文章"""
    api_client.test_get_posts()

def test_07_get_post(api_client):
    """测试获取单篇文章"""
    api_client.test_get_post()

def test_08_update_post(api_client):
    """测试更新文章"""
    api_client.test_update_post()

def test_09_create_comment(api_client):
    """测试创建评论"""
    api_client.test_create_comment()

def test_10_get_comments(api_client):
    """测试获取评论"""
    api_client.test_get_comments()

def test_11_delete_post(api_client):
    """测试删除文章"""
    api_client.test_delete_post()
