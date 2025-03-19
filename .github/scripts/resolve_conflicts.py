import os
import openai
import subprocess

# 读取环境变量
api_key = os.getenv("OPENAI_API_KEY")
base_url = os.getenv("OPENAI_BASE_URL", "https://api.openai.com/v1")  # 允许自定义 API 地址

# 配置 OpenAI
openai.api_key = api_key
openai.base_url = base_url  # 支持自定义 API 代理

def get_conflicted_files():
    """ 获取所有冲突文件的列表 """
    result = subprocess.run(["git", "diff", "--name-only", "--diff-filter=U"], capture_output=True, text=True)
    return result.stdout.strip().split("\n")

def get_conflict_content(file_path):
    """ 获取文件中的冲突内容 """
    with open(file_path, "r") as file:
        content = file.read()
    return content

def resolve_conflict_with_ai(content):
    """ 通过 OpenAI 解决冲突 """
    prompt = f"""
    你是一个高级软件工程师，我将提供一个代码文件，其中包含 Git 合并冲突 (`<<<<<<<`, `=======`, `>>>>>>>`)。
    请分析代码并自动解决冲突，提供一个最佳的合并版本。

    代码如下：
    ```
    {content}
    ```
    请输出合并后的完整代码，不要附加任何解释。
    """

    response = openai.ChatCompletion.create(
        model="gpt-4o",
        messages=[{"role": "user", "content": prompt}]
    )
    
    return response["choices"][0]["message"]["content"]

def apply_fix(file_path, fixed_content):
    """ 将 AI 解决的代码写回文件 """
    with open(file_path, "w") as file:
        file.write(fixed_content)

def main():
    """ 解决所有冲突的主逻辑 """
    conflicted_files = get_conflicted_files()
    
    for file_path in conflicted_files:
        if not file_path.strip():
            continue
        
        print(f"Resolving conflicts in: {file_path}")
        conflict_content = get_conflict_content(file_path)
        fixed_content = resolve_conflict_with_ai(conflict_content)
        apply_fix(file_path, fixed_content)
    
    print("All conflicts resolved.")

if __name__ == "__main__":
    main()
