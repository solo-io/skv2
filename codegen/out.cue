package v1

#CueBugSpec: {

	#ChildMessage: {
		name?: string @protobuf(1,string)
	}

	#ParentMessage: {
		myChild?: #ChildMessage @protobuf(2,ChildMessage,name=my_child)
	}
	{} | {
		myParent: #ParentMessage @protobuf(1,ParentMessage,name=my_parent)
	}
}

#CueBugStatus: {
	name?: string @protobuf(1,string)
}
