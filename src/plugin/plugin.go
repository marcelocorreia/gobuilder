package plugin

type Builder interface {
	Build() error
	Deploy() error
}

type Deployer interface {

}

type Installer interface {

}

type Documentation interface {

}

type Tester interface {

}




