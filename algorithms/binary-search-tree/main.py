class Node:
    def __init__(self, value, parent=None, left=None, right=None):
        self.value = value
        self.parent = parent
        self.left = left
        self.right = right

    def __repr__(self):
        return str(self.value)
        

class BinarySearchTree:
    def __init__(self, value):
        # todo: refactoring?
        # value argument can be None
        self.root = Node(value)

    def insert(self, value, root=None):
        # todo: refactor
        if root is None:
            self._insert(value, self.root)
        else:
            self._insert(value, root)
            
    def _insert(self, value, node):
        # refactor: can reduce nesting
        if value <= node.value:
            if node.left is None:
                new_node = Node(value, node)
                node.left = new_node
            else:
                self._insert(value, node.left)

        if value >= node.value:
            if node.right is None:
                new_node = Node(value, node)
                node.right = new_node
            else:
                self._insert(value, node.right)

    def search(self, value):
        if self.root is None:
            print('BST is empty')
            return

        return self._search(value, self.root)

    def _search(self, value, node):
        if node is None:
            return None
        
        if node.value == value:
            return node

        if value > node.value:
            return self._search(value, node.right)
        else:
            return self._search(value, node.left)

    def inorder_tree_walk(self, node=None):
        # todo: root does not exist (bst is empty)
        self._inorder_tree_walk(self.root)

    def _inorder_tree_walk(self, node):
        if node is not None:
            self._inorder_tree_walk(node.left)
            print('{} '.format(node.value))
            self._inorder_tree_walk(node.right)
    
    def preorder_tree_walk(self):
        # todo: root does not exist (bst is empty)
        self._preorder_tree_walk(self.root)

    def _preorder_tree_walk(self, node):
        if node is not None:
            print(node.value)
            self._preorder_tree_walk(node.left)
            self._preorder_tree_walk(node.right)

    def max_depth(self):
        if self.root is None:
            return 0
        return self._max_depth(self.root)

    def _max_depth(self, node):
        if node is None:
            return 0

        left_max_depth = self._max_depth(node.left)
        right_max_depth = self._max_depth(node.right)

        if left_max_depth > right_max_depth:
            return left_max_depth + 1
        
        return right_max_depth + 1

    def print_tree_2D(self):
        BASIC_SPACE = 10

        if self.root is None:
            print('Binary search tree is empry!')
            return

        self._print_tree_2D(self.root)

    def _print_tree_2D(self, node, space=0, BASIC_SPACE=10):
        if node is None:
            return

        space += BASIC_SPACE

        # Process right child
        self._print_tree_2D(node.right, space)

        # Print current node after space  
        # count  
        print()  
        for i in range(BASIC_SPACE, space): 
            print(end = " ")  
        print(node.value)  
    
        # Process left child  
        self._print_tree_2D(node.left, space) 


#    5
#   / \
#  9   7
# / |   | \
#10  5    9   7

# testing
bst = BinarySearchTree(5)
bst.insert(13)
bst.insert(4)
bst.insert(6)
bst.insert(12)

bst.insert(9)

root = bst.root

print('bst: ', bst)
print('root: ', root.value)
print('root.left: ', root.left.value)
print('root.right: ', root.right.value)

print('---')

bst.inorder_tree_walk()
print('@@@')
# bst.preorder_tree_walk()
depth = bst.max_depth()

print('depth: ', depth)

print('--------------------------')

bst.print_tree_2D()

print('------------------')

result = bst.search(12)

print('search result: ', result.value, result.left, result.right)
